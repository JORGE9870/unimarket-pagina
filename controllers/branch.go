package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type BranchController struct {
	BaseController
}

type Branch struct {
	Id         int64  `json:"id_sucursal"`
	BusinessId int64  `json:"id_negocio"`
	Name       string `json:"nombre"`
	Address    string `json:"direccion"`
	Phone      string `json:"telefono"`
	Status     string `json:"estado"`
	CreateDate string `json:"fecha_creacion"`
}

func (b *BranchController) Create() {
	var branch Branch
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &branch); err != nil {
		b.ResponseError("Datos inválidos", 400)
		return
	}

	if branch.Name == "" {
		b.ResponseError("El nombre es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&branch); err != nil {
		b.ResponseError("Error al crear sucursal", 500)
		return
	}

	b.ResponseSuccess(branch)
}

func (b *BranchController) GetAll() {
	var branches []Branch
	o := orm.NewOrm()
	qs := o.QueryTable("sucursales")

	if businessId, err := b.GetInt64("business_id"); err == nil {
		qs = qs.Filter("id_negocio", businessId)
	}
	if status := b.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}

	if _, err := qs.All(&branches); err != nil {
		b.ResponseError("Error al obtener sucursales", 500)
		return
	}

	b.ResponseSuccess(branches)
}

func (b *BranchController) Get() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	var branch Branch
	o := orm.NewOrm()
	if err := o.QueryTable("sucursales").Filter("id_sucursal", id).One(&branch); err != nil {
		b.ResponseError("Sucursal no encontrada", 404)
		return
	}

	b.ResponseSuccess(branch)
}

func (b *BranchController) Update() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	var branch Branch
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &branch); err != nil {
		b.ResponseError("Datos inválidos", 400)
		return
	}

	branch.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&branch); err != nil {
		b.ResponseError("Error al actualizar sucursal", 500)
		return
	}

	b.ResponseSuccess(branch)
}

func (b *BranchController) Delete() {
	id, err := b.GetInt64(":id")
	if err != nil {
		b.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Branch{Id: id}); err != nil {
		b.ResponseError("Error al eliminar sucursal", 500)
		return
	}

	b.ResponseSuccess(map[string]string{"message": "Sucursal eliminada"})
}
