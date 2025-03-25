package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type ProductController struct {
	BaseController
}

type Product struct {
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
func (p *ProductController) Create() {
	var product Product
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &product); err != nil {
		p.ResponseError("Datos inválidos", 400)
		return
	}

	if product.Name == "" {
		p.ResponseError("El nombre es requerido", 400)
		return
	}
	if product.BasePrice <= 0 {
		p.ResponseError("El precio base debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		p.ResponseError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Insert(&product); err != nil {
		tx.Rollback()
		p.ResponseError("Error al crear producto", 500)
		return
	}

	// Insertar categorías
	for _, catId := range product.Categories {
		if _, err := tx.Insert(&ProductCategory{
			ProductId:  product.Id,
			CategoryId: catId,
		}); err != nil {
			tx.Rollback()
			p.ResponseError("Error al asignar categorías", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		p.ResponseError("Error al confirmar transacción", 500)
		return
	}

	p.ResponseSuccess(product)
}

// @Title GetAllProducts
// @Description get all products
// @Success 200 {object} []models.Product
// @router / [get]
func (p *ProductController) GetAll() {
	var products []Product
	o := orm.NewOrm()
	qs := o.QueryTable("productos")

	// Filtros opcionales
	if businessId, err := p.GetInt64("business_id"); err == nil {
		qs = qs.Filter("id_negocio", businessId)
	}
	if status := p.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}
	if categoryId, err := p.GetInt("category_id"); err == nil {
		qs = qs.Filter("Categories__CategoryId", categoryId)
	}

	if _, err := qs.All(&products); err != nil {
		p.ResponseError("Error al obtener productos", 500)
		return
	}

	p.ResponseSuccess(products)
}

// @Title GetProduct
// @Description get product by id
// @Success 200 {object} models.Product
// @router /:id [get]
func (p *ProductController) Get() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	var product Product
	o := orm.NewOrm()
	if err := o.QueryTable("productos").Filter("id_producto", id).One(&product); err != nil {
		p.ResponseError("Producto no encontrado", 404)
		return
	}

	// Cargar categorías
	var categories []ProductCategory
	if _, err := o.QueryTable("producto_categorias").Filter("id_producto", id).All(&categories); err == nil {
		for _, cat := range categories {
			product.Categories = append(product.Categories, cat.CategoryId)
		}
	}

	p.ResponseSuccess(product)
}

// @Title UpdateProduct
// @Description update product
// @Success 200 {object} models.Product
// @router /:id [put]
func (p *ProductController) Update() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	var product Product
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &product); err != nil {
		p.ResponseError("Datos inválidos", 400)
		return
	}

	product.Id = id
	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		p.ResponseError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Update(&product); err != nil {
		tx.Rollback()
		p.ResponseError("Error al actualizar producto", 500)
		return
	}

	// Actualizar categorías
	if _, err := tx.QueryTable("producto_categorias").Filter("id_producto", id).Delete(); err != nil {
		tx.Rollback()
		p.ResponseError("Error al actualizar categorías", 500)
		return
	}

	for _, catId := range product.Categories {
		if _, err := tx.Insert(&ProductCategory{
			ProductId:  id,
			CategoryId: catId,
		}); err != nil {
			tx.Rollback()
			p.ResponseError("Error al asignar categorías", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		p.ResponseError("Error al confirmar transacción", 500)
		return
	}

	p.ResponseSuccess(product)
}

// @Title DeleteProduct
// @Description delete product
// @Success 200 {object} models.Product
// @router /:id [delete]
func (p *ProductController) Delete() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		p.ResponseError("Error al iniciar transacción", 500)
		return
	}

	// Eliminar categorías asociadas
	if _, err := tx.QueryTable("producto_categorias").Filter("id_producto", id).Delete(); err != nil {
		tx.Rollback()
		p.ResponseError("Error al eliminar categorías", 500)
		return
	}

	if _, err := tx.Delete(&Product{Id: id}); err != nil {
		tx.Rollback()
		p.ResponseError("Error al eliminar producto", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		p.ResponseError("Error al confirmar transacción", 500)
		return
	}

	p.ResponseSuccess(map[string]string{"message": "Producto eliminado"})
}
