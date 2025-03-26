package controllers

import (
	"encoding/json"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type CategoryController struct {
	BaseController
}

// @Title CreateCategory
// @Description create new category
// @Success 200 {object} models.Category
// @router / [post]
func (c *CategoryController) Create() {
	var category models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
		c.ResponseError("Datos de solicitud inválidos", 400)
		return
	}

	// Validaciones básicas
	if category.Name == "" {
		c.ResponseError("El nombre de la categoría es requerido", 400)
		return
	}

	oo := orm.NewOrm()

	// Verificar si ya existe una categoría con el mismo nombre
	var existingCategory models.Category
	if err := oo.QueryTable(new(models.Category)).Filter("Name", category.Name).One(&existingCategory); err == nil {
		c.ResponseError("Ya existe una categoría con ese nombre", 400)
		return
	}

	// Insertar la categoría
	if _, err := oo.Insert(&category); err != nil {
		c.ResponseError("Error al crear la categoría", 500)
		return
	}

	c.ResponseSuccess(category)
}

// @Title GetAllCategories
// @Description get all categories
// @Success 200 {object} []models.Category
// @router / [get]
func (c *CategoryController) GetAll() {
	var categories []models.Category
	oo := orm.NewOrm()

	if _, err := oo.QueryTable(new(models.Category)).All(&categories); err != nil {
		c.ResponseError("Error al obtener las categorías", 500)
		return
	}

	// Cargar productos relacionados
	for i := range categories {
		oo.LoadRelated(&categories[i], "Products")
	}

	c.ResponseSuccess(categories)
}

// @Title GetCategory
// @Description get category by id
// @Success 200 {object} models.Category
// @router /:id [get]
func (c *CategoryController) Get() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	var category models.Category
	oo := orm.NewOrm()
	if err := oo.QueryTable(new(models.Category)).Filter("Id", id).One(&category); err != nil {
		c.ResponseError("Categoría no encontrada", 404)
		return
	}

	// Cargar productos relacionados
	oo.LoadRelated(&category, "Products")

	c.ResponseSuccess(category)
}

// @Title UpdateCategory
// @Description update category
// @Success 200 {object} models.Category
// @router /:id [put]
func (c *CategoryController) Update() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	var category models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
		c.ResponseError("Datos de solicitud inválidos", 400)
		return
	}

	// Validaciones básicas
	if category.Name == "" {
		c.ResponseError("El nombre de la categoría es requerido", 400)
		return
	}

	category.Id = id
	oo := orm.NewOrm()

	// Verificar si ya existe otra categoría con el mismo nombre
	var existingCategory models.Category
	if err := oo.QueryTable(new(models.Category)).Filter("Name", category.Name).Exclude("Id", id).One(&existingCategory); err == nil {
		c.ResponseError("Ya existe una categoría con ese nombre", 400)
		return
	}

	// Actualizar la categoría
	if _, err := oo.Update(&category); err != nil {
		c.ResponseError("Error al actualizar la categoría", 500)
		return
	}

	c.ResponseSuccess(category)
}

// @Title DeleteCategory
// @Description delete category
// @Success 200 {object} models.Category
// @router /:id [delete]
func (c *CategoryController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	oo := orm.NewOrm()

	// Verificar si la categoría tiene productos asociados
	var count int64
	if count, err = oo.QueryTable(new(models.Product)).Filter("Category__Id", id).Count(); err != nil {
		c.ResponseError("Error al verificar productos asociados", 500)
		return
	}

	if count > 0 {
		c.ResponseError("No se puede eliminar la categoría porque tiene productos asociados", 400)
		return
	}

	// Eliminar la categoría
	if _, err := oo.Delete(&models.Category{BaseModel: models.BaseModel{Id: id}}); err != nil {
		c.ResponseError("Error al eliminar la categoría", 500)
		return
	}

	c.ResponseSuccess(map[string]string{"message": "Categoría eliminada exitosamente"})
}
