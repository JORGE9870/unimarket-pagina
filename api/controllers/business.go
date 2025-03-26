package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type BusinessController struct {
	BaseController
}

type Business struct {
	Id           int64  `json:"id_negocio"`
	Name         string `json:"nombre"`
	Description  string `json:"descripcion"`
	OwnerId      int64  `json:"id_propietario"`
	CategoryId   int    `json:"id_categoria"`
	LogoUrl      string `json:"logo_url"`
	Status       string `json:"estado"`
	RegisterDate string `json:"fecha_registro"`
}

func (b *BusinessController) Create() {
	var business Business
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &business); err != nil {
		b.ResponseError("Datos inválidos", 400)
		return
	}

	if business.Name == "" {
		b.ResponseError("El nombre es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&business); err != nil {
		b.ResponseError("Error al crear negocio", 500)
		return
	}

	b.ResponseSuccess(business)
}

func (b *BusinessController) GetAll() {
	var businesses []Business
	o := orm.NewOrm()
	qs := o.QueryTable("negocios")

	// Filtros opcionales
	if ownerId, err := b.GetInt64("owner_id"); err == nil {
		qs = qs.Filter("id_propietario", ownerId)
	}
	if categoryId, err := b.GetInt("category_id"); err == nil {
		qs = qs.Filter("id_categoria", categoryId)
	}
	if status := b.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}

	if _, err := qs.All(&businesses); err != nil {
		b.ResponseError("Error al obtener negocios", 500)
		return
	}

	b.ResponseSuccess(businesses)
}

func (b *BusinessController) Get() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	var business Business
	o := orm.NewOrm()
	if err := o.QueryTable("negocios").Filter("id_negocio", id).One(&business); err != nil {
		b.ResponseError("Negocio no encontrado", 404)
		return
	}

	b.ResponseSuccess(business)
}

func (b *BusinessController) Update() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	var business Business
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &business); err != nil {
		b.ResponseError("Datos inválidos", 400)
		return
	}

	business.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&business); err != nil {
		b.ResponseError("Error al actualizar negocio", 500)
		return
	}

	b.ResponseSuccess(business)
}

func (b *BusinessController) Delete() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()

	// Verificar si hay sucursales asociadas
	count, err := o.QueryTable("sucursales").Filter("id_negocio", id).Count()
	if err != nil {
		b.ResponseError("Error al verificar sucursales", 500)
		return
	}
	if count > 0 {
		b.ResponseError("No se puede eliminar el negocio porque tiene sucursales asociadas", 400)
		return
	}

	if _, err := o.Delete(&Business{Id: id}); err != nil {
		b.ResponseError("Error al eliminar negocio", 500)
		return
	}

	b.ResponseSuccess(map[string]string{"message": "Negocio eliminado"})
}
