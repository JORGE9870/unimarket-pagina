package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorPago struct {
	ControladorBase
}

type Pago struct {
	Id         int64   `json:"id_pago"`
	PedidoId   int64   `json:"id_pedido"`
	Monto      float64 `json:"monto"`
	Estado     string  `json:"estado"`
	CreateDate string  `json:"fecha_creacion"`
}

func (c *ControladorPago) Crear() {
	var pago Pago
	if err := c.ParsearYValidarJSON(&pago); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if pago.PedidoId == 0 {
		c.RespuestaError("El pedido es requerido", 400)
		return
	}
	if pago.Monto <= 0 {
		c.RespuestaError("El monto debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&pago); err != nil {
		c.RespuestaError("Error al crear pago", 500)
		return
	}

	c.RespuestaExito(pago)
}

func (c *ControladorPago) Listar() {
	var pagos []Pago
	o := orm.NewOrm()
	qs := o.QueryTable("pagos")

	if pedidoId, err := c.GetInt64("pedido_id"); err == nil {
		qs = qs.Filter("id_pedido", pedidoId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&pagos); err != nil {
		c.RespuestaError("Error al obtener pagos", 500)
		return
	}

	c.RespuestaExito(pagos)
}

func (c *ControladorPago) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var pago Pago
	o := orm.NewOrm()
	if err := o.QueryTable("pagos").Filter("id_pago", id).One(&pago); err != nil {
		c.RespuestaError("Pago no encontrado", 404)
		return
	}

	c.RespuestaExito(pago)
}

func (c *ControladorPago) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var pago Pago
	if err := c.ParsearYValidarJSON(&pago); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if pago.Monto <= 0 {
		c.RespuestaError("El monto debe ser mayor a 0", 400)
		return
	}

	pago.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&pago); err != nil {
		c.RespuestaError("Error al actualizar pago", 500)
		return
	}

	c.RespuestaExito(pago)
}

func (c *ControladorPago) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Pago{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar pago", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Pago eliminado"})
}
