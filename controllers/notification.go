package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type NotificationController struct {
	BaseController
}

type Notification struct {
	Id         int64  `json:"id_notificacion"`
	UserId     int64  `json:"id_usuario"`
	Title      string `json:"titulo"`
	Message    string `json:"mensaje"`
	Type       string `json:"tipo"`
	IsRead     bool   `json:"leida"`
	CreateDate string `json:"fecha_creacion"`
}

func (n *NotificationController) Create() {
	var notification Notification
	if err := json.Unmarshal(n.Ctx.Input.RequestBody, &notification); err != nil {
		n.ResponseError("Datos inválidos", 400)
		return
	}

	if notification.UserId == 0 {
		n.ResponseError("El usuario es requerido", 400)
		return
	}
	if notification.Title == "" {
		n.ResponseError("El título es requerido", 400)
		return
	}
	if notification.Message == "" {
		n.ResponseError("El mensaje es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&notification); err != nil {
		n.ResponseError("Error al crear notificación", 500)
		return
	}

	n.ResponseSuccess(notification)
}

func (n *NotificationController) GetAll() {
	var notifications []Notification
	o := orm.NewOrm()
	qs := o.QueryTable("notificaciones")

	if userId, err := n.GetInt64("user_id"); err == nil {
		qs = qs.Filter("id_usuario", userId)
	}
	if isRead, err := n.GetBool("is_read"); err == nil {
		qs = qs.Filter("leida", isRead)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&notifications); err != nil {
		n.ResponseError("Error al obtener notificaciones", 500)
		return
	}

	n.ResponseSuccess(notifications)
}

func (n *NotificationController) Get() {
	id, err := n.GetInt64(":id")
	if err != nil {
		n.ResponseError("ID inválido", 400)
		return
	}

	var notification Notification
	o := orm.NewOrm()
	if err := o.QueryTable("notificaciones").Filter("id_notificacion", id).One(&notification); err != nil {
		n.ResponseError("Notificación no encontrada", 404)
		return
	}

	n.ResponseSuccess(notification)
}

func (n *NotificationController) MarkAsRead() {
	id, err := n.GetInt64(":id")
	if err != nil {
		n.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.QueryTable("notificaciones").Filter("id_notificacion", id).Update(orm.Params{
		"leida": true,
	}); err != nil {
		n.ResponseError("Error al marcar notificación como leída", 500)
		return
	}

	n.ResponseSuccess(map[string]string{"message": "Notificación marcada como leída"})
}

func (n *NotificationController) MarkAllAsRead() {
	userId, err := n.GetInt64("user_id")
	if err != nil {
		n.ResponseError("ID de usuario inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.QueryTable("notificaciones").Filter("id_usuario", userId).Update(orm.Params{
		"leida": true,
	}); err != nil {
		n.ResponseError("Error al marcar notificaciones como leídas", 500)
		return
	}

	n.ResponseSuccess(map[string]string{"message": "Todas las notificaciones marcadas como leídas"})
}

func (n *NotificationController) Delete() {
	id, err := n.GetInt64(":id")
	if err != nil {
		n.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Notification{Id: id}); err != nil {
		n.ResponseError("Error al eliminar notificación", 500)
		return
	}

	n.ResponseSuccess(map[string]string{"message": "Notificación eliminada"})
}
