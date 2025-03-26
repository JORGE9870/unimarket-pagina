package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type PaymentController struct {
	BaseController
}

type Payment struct {
	Id         int64   `json:"id_pago"`
	OrderId    int64   `json:"id_pedido"`
	Amount     float64 `json:"monto"`
	Method     string  `json:"metodo"`
	Status     string  `json:"estado"`
	Reference  string  `json:"referencia"`
	CreateDate string  `json:"fecha_creacion"`
}

func (p *PaymentController) Create() {
	var payment Payment
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &payment); err != nil {
		p.ResponseError("Datos inválidos", 400)
		return
	}

	if payment.OrderId == 0 {
		p.ResponseError("El pedido es requerido", 400)
		return
	}
	if payment.Amount <= 0 {
		p.ResponseError("El monto debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&payment); err != nil {
		p.ResponseError("Error al crear pago", 500)
		return
	}

	p.ResponseSuccess(payment)
}

func (p *PaymentController) GetAll() {
	var payments []Payment
	o := orm.NewOrm()
	qs := o.QueryTable("pagos")

	if orderId, err := p.GetInt64("order_id"); err == nil {
		qs = qs.Filter("id_pedido", orderId)
	}
	if status := p.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}

	if _, err := qs.All(&payments); err != nil {
		p.ResponseError("Error al obtener pagos", 500)
		return
	}

	p.ResponseSuccess(payments)
}

func (p *PaymentController) Get() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	var payment Payment
	o := orm.NewOrm()
	if err := o.QueryTable("pagos").Filter("id_pago", id).One(&payment); err != nil {
		p.ResponseError("Pago no encontrado", 404)
		return
	}

	p.ResponseSuccess(payment)
}

func (p *PaymentController) Update() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	var payment Payment
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &payment); err != nil {
		p.ResponseError("Datos inválidos", 400)
		return
	}

	payment.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&payment); err != nil {
		p.ResponseError("Error al actualizar pago", 500)
		return
	}

	p.ResponseSuccess(payment)
}

func (p *PaymentController) Delete() {
	id, err := p.GetInt64(":id")
	if err != nil {
		p.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Payment{Id: id}); err != nil {
		p.ResponseError("Error al eliminar pago", 500)
		return
	}

	p.ResponseSuccess(map[string]string{"message": "Pago eliminado"})
}
