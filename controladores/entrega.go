package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorEntrega struct {
	ControladorBase
}

type Entrega struct {
	Id         int64  `json:"id_entrega"`
	PedidoId   int64  `json:"id_pedido"`
	Estado     string `json:"estado"`
	Direccion  string `json:"direccion"`
	CreateDate string `json:"fecha_creacion"`
}

func (c *ControladorEntrega) Crear() {
	var entrega Entrega
	if err := c.ParsearYValidarJSON(&entrega); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if entrega.PedidoId == 0 {
		c.RespuestaError("El pedido es requerido", 400)
		return
	}
	if entrega.Direccion == "" {
		c.RespuestaError("La dirección es requerida", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&entrega); err != nil {
		c.RespuestaError("Error al crear entrega", 500)
		return
	}

	c.RespuestaExito(entrega)
}

func (c *ControladorEntrega) Listar() {
	var entregas []Entrega
	o := orm.NewOrm()
	qs := o.QueryTable("entregas")

	if pedidoId, err := c.GetInt64("pedido_id"); err == nil {
		qs = qs.Filter("id_pedido", pedidoId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&entregas); err != nil {
		c.RespuestaError("Error al obtener entregas", 500)
		return
	}

	c.RespuestaExito(entregas)
}

func (c *ControladorEntrega) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var entrega Entrega
	o := orm.NewOrm()
	if err := o.QueryTable("entregas").Filter("id_entrega", id).One(&entrega); err != nil {
		c.RespuestaError("Entrega no encontrada", 404)
		return
	}

	c.RespuestaExito(entrega)
}

func (c *ControladorEntrega) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var entrega Entrega
	if err := c.ParsearYValidarJSON(&entrega); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if entrega.Direccion == "" {
		c.RespuestaError("La dirección es requerida", 400)
		return
	}

	entrega.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&entrega); err != nil {
		c.RespuestaError("Error al actualizar entrega", 500)
		return
	}

	c.RespuestaExito(entrega)
}

func (c *ControladorEntrega) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Entrega{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar entrega", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Entrega eliminada"})
}
