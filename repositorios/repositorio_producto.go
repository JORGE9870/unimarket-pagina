package repositories

import (
	"context"
	"errors"
	"time"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type ProductRepository struct {
	orm orm.Ormer
	ctx context.Context
}

func (r *ProductRepository) GetProductDetails(productId string) (*models.ProductDetails, error) {
	var details models.ProductDetails

	// Consulta principal con joins
	sql := `
        SELECT 
            p.*,
            c.nombre as categoria_nombre,
            m.nombre as marca_nombre,
            AVG(r.puntuacion) as rating_promedio,
            COUNT(DISTINCT r.id) as total_reviews,
            COUNT(DISTINCT v.id) as total_ventas
        FROM productos p
        LEFT JOIN categorias c ON p.id_categoria = c.id
        LEFT JOIN marcas m ON p.id_marca = m.id
        LEFT JOIN reviews r ON p.id = r.id_producto
        LEFT JOIN ventas v ON p.id = v.id_producto
        WHERE p.id = ?
        GROUP BY p.id
    `

	err := r.orm.Raw(sql, productId).QueryRow(&details)
	if err != nil {
		return nil, err
	}

	// Cargar ubicaciones disponibles
	sql = `
        SELECT 
            s.id,
            s.nombre,
            si.cantidad
        FROM stock_inventario si
        JOIN sucursales s ON si.id_sucursal = s.id
        WHERE si.id_producto = ?
        AND si.cantidad > 0
    `

	_, err = r.orm.Raw(sql, productId).QueryRows(&details.Locations)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func (r *ProductRepository) SearchProducts(filters models.ProductFilters) ([]*models.Product, error) {
	qs := r.orm.QueryTable("productos")

	// Aplicar filtros
	if filters.Category != "" {
		qs = qs.Filter("id_categoria", filters.Category)
	}
	if filters.Brand != "" {
		qs = qs.Filter("id_marca", filters.Brand)
	}
	if filters.MinPrice > 0 {
		qs = qs.Filter("precio__gte", filters.MinPrice)
	}
	if filters.MaxPrice > 0 {
		qs = qs.Filter("precio__lte", filters.MaxPrice)
	}
	if filters.Status != "" {
		qs = qs.Filter("estado", filters.Status)
	}

	// Ordenamiento
	if filters.SortBy != "" {
		if filters.SortDesc {
			qs = qs.OrderBy("-" + filters.SortBy)
		} else {
			qs = qs.OrderBy(filters.SortBy)
		}
	}

	// Paginación
	if filters.Page > 0 && filters.PageSize > 0 {
		offset := (filters.Page - 1) * filters.PageSize
		qs = qs.Limit(filters.PageSize).Offset(offset)
	}

	var products []*models.Product
	_, err := qs.All(&products)
	return products, err
}

func (r *ProductRepository) UpdateProductStock(productId string, quantity int, operation string) error {
	tx, err := r.orm.Begin()
	if err != nil {
		return err
	}

	sql := ""
	switch operation {
	case "increment":
		sql = "UPDATE productos SET stock = stock + ? WHERE id = ?"
	case "decrement":
		sql = "UPDATE productos SET stock = stock - ? WHERE id = ? AND stock >= ?"
	default:
		tx.Rollback()
		return errors.New("operación no válida")
	}

	result, err := tx.Raw(sql, quantity, productId, quantity).Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		tx.Rollback()
		return errors.New("stock insuficiente o producto no encontrado")
	}

	// Registrar movimiento de inventario
	movement := &models.StockMovement{
		ProductID: productId,
		Quantity:  quantity,
		Operation: operation,
		Timestamp: time.Now(),
	}

	_, err = tx.Insert(movement)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
