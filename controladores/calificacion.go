package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorCalificacion struct {
	ControladorBase
}

type Calificacion struct {
	Id         int64   `json:"id_valoracion"`
	ProductId  int64   `json:"id_producto"`
	UserId     int64   `json:"id_usuario"`
	Score      float64 `json:"puntuacion"`
	Comment    string  `json:"comentario"`
	CreateDate string  `json:"fecha_creacion"`
}

func (r *ControladorCalificacion) Crear() {
	var calificacion Calificacion
	if err := r.ParsearYValidarJSON(&calificacion); err != nil {
		r.RespuestaError("Datos inválidos", 400)
		return
	}

	if calificacion.ProductId == 0 {
		r.RespuestaError("El producto es requerido", 400)
		return
	}
	if calificacion.UserId == 0 {
		r.RespuestaError("El usuario es requerido", 400)
		return
	}
	if calificacion.Score < 1 || calificacion.Score > 5 {
		r.RespuestaError("La puntuación debe estar entre 1 y 5", 400)
		return
	}

	o := orm.NewOrm()

	var calificacionExistente Calificacion
	err := o.QueryTable("valoraciones").
		Filter("id_producto", calificacion.ProductId).
		Filter("id_usuario", calificacion.UserId).
		One(&calificacionExistente)

	if err == nil {
		r.RespuestaError("Ya has valorado este producto", 400)
		return
	}

	if _, err := o.Insert(&calificacion); err != nil {
		r.RespuestaError("Error al crear valoración", 500)
		return
	}

	r.RespuestaExito(calificacion)
}

func (r *ControladorCalificacion) Listar() {
	var calificaciones []Calificacion
	o := orm.NewOrm()
	qs := o.QueryTable("valoraciones")

	if productId, err := r.GetInt64("product_id"); err == nil {
		qs = qs.Filter("id_producto", productId)
	}
	if userId, err := r.GetInt64("user_id"); err == nil {
		qs = qs.Filter("id_usuario", userId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&calificaciones); err != nil {
		r.RespuestaError("Error al obtener valoraciones", 500)
		return
	}

	r.RespuestaExito(calificaciones)
}

func (r *ControladorCalificacion) Obtener() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.RespuestaError("ID inválido", 400)
		return
	}

	var calificacion Calificacion
	o := orm.NewOrm()
	if err := o.QueryTable("valoraciones").Filter("id_valoracion", id).One(&calificacion); err != nil {
		r.RespuestaError("Valoración no encontrada", 404)
		return
	}

	r.RespuestaExito(calificacion)
}

func (r *ControladorCalificacion) Actualizar() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.RespuestaError("ID inválido", 400)
		return
	}

	var calificacion Calificacion
	if err := r.ParsearYValidarJSON(&calificacion); err != nil {
		r.RespuestaError("Datos inválidos", 400)
		return
	}

	if calificacion.Score < 1 || calificacion.Score > 5 {
		r.RespuestaError("La puntuación debe estar entre 1 y 5", 400)
		return
	}

	calificacion.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&calificacion); err != nil {
		r.RespuestaError("Error al actualizar valoración", 500)
		return
	}

	r.RespuestaExito(calificacion)
}

func (r *ControladorCalificacion) Eliminar() {
	id, err := r.GetInt64(":id")
	if err != nil {
		r.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Calificacion{Id: id}); err != nil {
		r.RespuestaError("Error al eliminar valoración", 500)
		return
	}

	r.RespuestaExito(map[string]string{"mensaje": "Valoración eliminada"})
}
