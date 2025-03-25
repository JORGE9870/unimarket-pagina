package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type DeliveryController struct {
	BaseController
}

type Delivery struct {
	Id          int64   `json:"id_entrega"`
	OrderId     int64   `json:"id_pedido"`
	DeliveryMan int64   `json:"id_repartidor"`
	Status      string  `json:"estado"`
	Address     string  `json:"direccion"`
	Latitude    float64 `json:"latitud"`
	Longitude   float64 `json:"longitud"`
	StartTime   string  `json:"hora_inicio"`
	EndTime     string  `json:"hora_fin"`
	Notes       string  `json:"notas"`
}

func (d *DeliveryController) Create() {
	var delivery Delivery
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &delivery); err != nil {
		d.ResponseError("Datos inválidos", 400)
		return
	}

	if delivery.OrderId == 0 {
		d.ResponseError("El pedido es requerido", 400)
		return
	}
	if delivery.Address == "" {
		d.ResponseError("La dirección es requerida", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&delivery); err != nil {
		d.ResponseError("Error al crear entrega", 500)
		return
	}

	d.ResponseSuccess(delivery)
}

func (d *DeliveryController) GetAll() {
	var deliveries []Delivery
	o := orm.NewOrm()
	qs := o.QueryTable("entregas")

	if orderId, err := d.GetInt64("order_id"); err == nil {
		qs = qs.Filter("id_pedido", orderId)
	}
	if deliveryManId, err := d.GetInt64("delivery_man_id"); err == nil {
		qs = qs.Filter("id_repartidor", deliveryManId)
	}
	if status := d.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}

	if _, err := qs.All(&deliveries); err != nil {
		d.ResponseError("Error al obtener entregas", 500)
		return
	}

	d.ResponseSuccess(deliveries)
}

func (d *DeliveryController) Get() {
	id, err := d.GetInt64(":id")
	if err != nil {
		d.ResponseError("ID inválido", 400)
		return
	}

	var delivery Delivery
	o := orm.NewOrm()
	if err := o.QueryTable("entregas").Filter("id_entrega", id).One(&delivery); err != nil {
		d.ResponseError("Entrega no encontrada", 404)
		return
	}

	d.ResponseSuccess(delivery)
}

func (d *DeliveryController) Update() {
	id, err := d.GetInt64(":id")
	if err != nil {
		d.ResponseError("ID inválido", 400)
		return
	}

	var delivery Delivery
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &delivery); err != nil {
		d.ResponseError("Datos inválidos", 400)
		return
	}

	delivery.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&delivery); err != nil {
		d.ResponseError("Error al actualizar entrega", 500)
		return
	}

	d.ResponseSuccess(delivery)
}

func (d *DeliveryController) Delete() {
	id, err := d.GetInt64(":id")
	if err != nil {
		d.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Delivery{Id: id}); err != nil {
		d.ResponseError("Error al eliminar entrega", 500)
		return
	}

	d.ResponseSuccess(map[string]string{"message": "Entrega eliminada"})
}
