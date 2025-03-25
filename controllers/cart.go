package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type CartController struct {
	BaseController
}

type Cart struct {
	Id         int64      `json:"id_carrito"`
	UserId     int64      `json:"id_usuario"`
	BranchId   int64      `json:"id_sucursal"`
	Status     string     `json:"estado"`
	CreateDate string     `json:"fecha_creacion"`
	Items      []CartItem `json:"items,omitempty"`
}

type CartItem struct {
	Id        int64   `json:"id_item"`
	CartId    int64   `json:"id_carrito"`
	ProductId int64   `json:"id_producto"`
	Quantity  int     `json:"cantidad"`
	Price     float64 `json:"precio"`
}

func (c *CartController) Create() {
	var cart Cart
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &cart); err != nil {
		c.ResponseError("Datos inválidos", 400)
		return
	}

	if cart.UserId == 0 {
		c.ResponseError("El usuario es requerido", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.ResponseError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Insert(&cart); err != nil {
		tx.Rollback()
		c.ResponseError("Error al crear carrito", 500)
		return
	}

	for _, item := range cart.Items {
		item.CartId = cart.Id
		if _, err := tx.Insert(&item); err != nil {
			tx.Rollback()
			c.ResponseError("Error al agregar items al carrito", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.ResponseError("Error al confirmar transacción", 500)
		return
	}

	c.ResponseSuccess(cart)
}

func (c *CartController) GetAll() {
	var carts []Cart
	o := orm.NewOrm()
	qs := o.QueryTable("carritos")

	if userId, err := c.GetInt64("user_id"); err == nil {
		qs = qs.Filter("id_usuario", userId)
	}
	if status := c.GetString("status"); status != "" {
		qs = qs.Filter("estado", status)
	}

	if _, err := qs.All(&carts); err != nil {
		c.ResponseError("Error al obtener carritos", 500)
		return
	}

	for i := range carts {
		var items []CartItem
		if _, err := o.QueryTable("items_carrito").Filter("id_carrito", carts[i].Id).All(&items); err == nil {
			carts[i].Items = items
		}
	}

	c.ResponseSuccess(carts)
}

func (c *CartController) Get() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	var cart Cart
	o := orm.NewOrm()
	if err := o.QueryTable("carritos").Filter("id_carrito", id).One(&cart); err != nil {
		c.ResponseError("Carrito no encontrado", 404)
		return
	}

	var items []CartItem
	if _, err := o.QueryTable("items_carrito").Filter("id_carrito", id).All(&items); err == nil {
		cart.Items = items
	}

	c.ResponseSuccess(cart)
}

func (c *CartController) Update() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	var cart Cart
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &cart); err != nil {
		c.ResponseError("Datos inválidos", 400)
		return
	}

	cart.Id = id
	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.ResponseError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.Update(&cart); err != nil {
		tx.Rollback()
		c.ResponseError("Error al actualizar carrito", 500)
		return
	}

	if _, err := tx.QueryTable("items_carrito").Filter("id_carrito", id).Delete(); err != nil {
		tx.Rollback()
		c.ResponseError("Error al actualizar items", 500)
		return
	}

	for _, item := range cart.Items {
		item.CartId = id
		if _, err := tx.Insert(&item); err != nil {
			tx.Rollback()
			c.ResponseError("Error al actualizar items", 500)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.ResponseError("Error al confirmar transacción", 500)
		return
	}

	c.ResponseSuccess(cart)
}

func (c *CartController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.ResponseError("Error al iniciar transacción", 500)
		return
	}

	if _, err := tx.QueryTable("items_carrito").Filter("id_carrito", id).Delete(); err != nil {
		tx.Rollback()
		c.ResponseError("Error al eliminar items", 500)
		return
	}

	if _, err := tx.Delete(&Cart{Id: id}); err != nil {
		tx.Rollback()
		c.ResponseError("Error al eliminar carrito", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		c.ResponseError("Error al confirmar transacción", 500)
		return
	}

	c.ResponseSuccess(map[string]string{"message": "Carrito eliminado"})
}
