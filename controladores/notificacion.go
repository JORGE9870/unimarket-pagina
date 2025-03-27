package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorNotificacion struct {
	ControladorBase
}

type Notificacion struct {
	Id         int64  `json:"id_notificacion"`
	UsuarioId  int64  `json:"id_usuario"`
	Titulo     string `json:"titulo"`
	Mensaje    string `json:"mensaje"`
	Leida      bool   `json:"leida"`
	CreateDate string `json:"fecha_creacion"`
}

func (c *ControladorNotificacion) Crear() {
	var notificacion Notificacion
	if err := c.ParsearYValidarJSON(&notificacion); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if notificacion.UsuarioId == 0 {
		c.RespuestaError("El usuario es requerido", 400)
		return
	}
	if notificacion.Titulo == "" {
		c.RespuestaError("El título es requerido", 400)
		return
	}
	if notificacion.Mensaje == "" {
		c.RespuestaError("El mensaje es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&notificacion); err != nil {
		c.RespuestaError("Error al crear notificación", 500)
		return
	}

	c.RespuestaExito(notificacion)
}

func (c *ControladorNotificacion) Listar() {
	var notificaciones []Notificacion
	o := orm.NewOrm()
	qs := o.QueryTable("notificaciones")

	if usuarioId, err := c.GetInt64("usuario_id"); err == nil {
		qs = qs.Filter("id_usuario", usuarioId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&notificaciones); err != nil {
		c.RespuestaError("Error al obtener notificaciones", 500)
		return
	}

	c.RespuestaExito(notificaciones)
}

func (c *ControladorNotificacion) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var notificacion Notificacion
	o := orm.NewOrm()
	if err := o.QueryTable("notificaciones").Filter("id_notificacion", id).One(&notificacion); err != nil {
		c.RespuestaError("Notificación no encontrada", 404)
		return
	}

	c.RespuestaExito(notificacion)
}

func (c *ControladorNotificacion) MarcarLeida() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	_, err = o.Raw("UPDATE notificaciones SET leida = true WHERE id_notificacion = ?", id).Exec()
	if err != nil {
		c.RespuestaError("Error al marcar notificación como leída", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Notificación marcada como leída"})
}

func (c *ControladorNotificacion) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Notificacion{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar notificación", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Notificación eliminada"})
}
