package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorProducto struct {
	ControladorBase
}

type Producto struct {
	Id          int64   `json:"id_producto"`
	BusinessId  int64   `json:"id_negocio"`
	Name        string  `json:"nombre"`
	Description string  `json:"descripcion"`
	BasePrice   float64 `json:"precio_base"`
	Stock       int     `json:"stock"`
	SKU         string  `json:"sku"`
	Status      string  `json:"estado"`
	CreateDate  string  `json:"fecha_creacion"`
	Categories  []int   `json:"categorias,omitempty"`
}

type ProductCategory struct {
	ProductId  int64 `json:"id_producto"`
	CategoryId int   `json:"id_categoria"`
}

// @Title CreateProduct
// @Description create new product
// @Success 200 {object} models.Product
// @router / [post]
func (c *ControladorProducto) Crear() {
	var producto Producto
	if err := c.ParsearYValidarJSON(&producto); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if producto.Name == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}
	if producto.BasePrice <= 0 {
		c.RespuestaError("El precio base debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Insert(&producto); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al crear producto", 500)
		return
	}

	// Insertar categorías
	for _, catId := range producto.Categories {
		if _, err := tx.Insert(&ProductCategory{
			ProductId:  producto.Id,
			CategoryId: catId,
		}); err != nil {
			tx.Rollback()
			c.RespuestaError("Error al asignar categorías", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(producto)
}

// @Title GetAllProducts
// @Description get all products
// @Success 200 {object} []models.Product
// @router / [get]
func (c *ControladorProducto) Listar() {
	var productos []Producto
	o := orm.NewOrm()
	qs := o.QueryTable("productos")

	// Filtros opcionales
	if businessId, err := c.GetInt64("business_id"); err == nil {
		qs = qs.Filter("id_negocio", businessId)
	}
	if status := c.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}
	if categoryId, err := c.GetInt("category_id"); err == nil {
		qs = qs.Filter("Categories__CategoryId", categoryId)
	}

	if _, err := qs.All(&productos); err != nil {
		c.RespuestaError("Error al obtener productos", 500)
		return
	}

	c.RespuestaExito(productos)
}

// @Title GetProduct
// @Description get product by id
// @Success 200 {object} models.Product
// @router /:id [get]
func (c *ControladorProducto) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var producto Producto
	o := orm.NewOrm()
	if err := o.QueryTable("productos").Filter("id_producto", id).One(&producto); err != nil {
		c.RespuestaError("Producto no encontrado", 404)
		return
	}

	// Cargar categorías
	var categories []ProductCategory
	if _, err := o.QueryTable("producto_categorias").Filter("id_producto", id).All(&categories); err == nil {
		for _, cat := range categories {
			producto.Categories = append(producto.Categories, cat.CategoryId)
		}
	}

	c.RespuestaExito(producto)
}

// @Title UpdateProduct
// @Description update product
// @Success 200 {object} models.Product
// @router /:id [put]
func (c *ControladorProducto) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var producto Producto
	if err := c.ParsearYValidarJSON(&producto); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if producto.Name == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}
	if producto.BasePrice <= 0 {
		c.RespuestaError("El precio base debe ser mayor a 0", 400)
		return
	}

	producto.Id = id
	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Update(&producto); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al actualizar producto", 500)
		return
	}

	// Actualizar categorías
	if _, err := tx.QueryTable("producto_categorias").Filter("id_producto", id).Delete(); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al actualizar categorías", 500)
		return
	}

	for _, catId := range producto.Categories {
		if _, err := tx.Insert(&ProductCategory{
			ProductId:  id,
			CategoryId: catId,
		}); err != nil {
			tx.Rollback()
			c.RespuestaError("Error al asignar categorías", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(producto)
}

// @Title DeleteProduct
// @Description delete product
// @Success 200 {object} models.Product
// @router /:id [delete]
func (c *ControladorProducto) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	// Eliminar categorías asociadas
	if _, err := tx.QueryTable("producto_categorias").Filter("id_producto", id).Delete(); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al eliminar categorías", 500)
		return
	}

	if _, err := tx.Delete(&Producto{Id: id}); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al eliminar producto", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Producto eliminado"})
}
