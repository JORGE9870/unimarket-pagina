package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type ReviewController struct {
	BaseController
}

type Review struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id"`
	UserId    int64  `json:"user_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

func (r *ReviewController) Create() {
	var review Review
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &review); err != nil {
		r.ResponseError("Datos inválidos", 400)
		return
	}

	if review.Rating < 1 || review.Rating > 5 {
		r.ResponseError("La calificación debe estar entre 1 y 5", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&review); err != nil {
		r.ResponseError("Error al crear reseña", 500)
		return
	}

	r.ResponseSuccess(review)
}
