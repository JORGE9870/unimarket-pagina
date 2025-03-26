package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type RatingController struct {
	BaseController
}

type Rating struct {
	Id         int64   `json:"id_valoracion"`
	ProductId  int64   `json:"id_producto"`
	UserId     int64   `json:"id_usuario"`
	Score      float64 `json:"puntuacion"`
	Comment    string  `json:"comentario"`
	CreateDate string  `json:"fecha_creacion"`
}

func (r *RatingController) Create() {
	var rating Rating
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &rating); err != nil {
		r.ResponseError("Datos inválidos", 400)
		return
	}

	if rating.ProductId == 0 {
		r.ResponseError("El producto es requerido", 400)
		return
	}
	if rating.UserId == 0 {
		r.ResponseError("El usuario es requerido", 400)
		return
	}
	if rating.Score < 1 || rating.Score > 5 {
		r.ResponseError("La puntuación debe estar entre 1 y 5", 400)
		return
	}

	o := orm.NewOrm()

	// Verificar si el usuario ya valoró este producto
	var existingRating Rating
	err := o.QueryTable("valoraciones").
		Filter("id_producto", rating.ProductId).
		Filter("id_usuario", rating.UserId).
		One(&existingRating)

	if err == nil {
		r.ResponseError("Ya has valorado este producto", 400)
		return
	}

	if _, err := o.Insert(&rating); err != nil {
		r.ResponseError("Error al crear valoración", 500)
		return
	}

	r.ResponseSuccess(rating)
}

func (r *RatingController) GetAll() {
	var ratings []Rating
	o := orm.NewOrm()
	qs := o.QueryTable("valoraciones")

	if productId, err := r.GetInt64("product_id"); err == nil {
		qs = qs.Filter("id_producto", productId)
	}
	if userId, err := r.GetInt64("user_id"); err == nil {
		qs = qs.Filter("id_usuario", userId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&ratings); err != nil {
		r.ResponseError("Error al obtener valoraciones", 500)
		return
	}

	r.ResponseSuccess(ratings)
}

func (r *RatingController) Get() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.ResponseError("ID inválido", 400)
		return
	}

	var rating Rating
	o := orm.NewOrm()
	if err := o.QueryTable("valoraciones").Filter("id_valoracion", id).One(&rating); err != nil {
		r.ResponseError("Valoración no encontrada", 404)
		return
	}

	r.ResponseSuccess(rating)
}

func (r *RatingController) Update() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.ResponseError("ID inválido", 400)
		return
	}

	var rating Rating
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &rating); err != nil {
		r.ResponseError("Datos inválidos", 400)
		return
	}

	if rating.Score < 1 || rating.Score > 5 {
		r.ResponseError("La puntuación debe estar entre 1 y 5", 400)
		return
	}

	rating.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&rating); err != nil {
		r.ResponseError("Error al actualizar valoración", 500)
		return
	}

	r.ResponseSuccess(rating)
}

func (r *RatingController) Delete() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Rating{Id: id}); err != nil {
		r.ResponseError("Error al eliminar valoración", 500)
		return
	}

	r.ResponseSuccess(map[string]string{"message": "Valoración eliminada"})
}
