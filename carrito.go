package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorCarrito struct {
	ControladorBase
}

type Carrito struct {
	Id         int64  `json:"id_carrito"`
	UserId     int64  `json:"id_usuario"`
	ProductId  int64  `json:"id_producto"`
	Cantidad   int    `json:"cantidad"`
	CreateDate string `json:"fecha_creacion"`
}

func (c *ControladorCarrito) Crear() {
	var carrito Carrito
	if err := c.ParsearYValidarJSON(&carrito); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if carrito.ProductId == 0 {
		c.RespuestaError("El producto es requerido", 400)
		return
	}
	if carrito.Cantidad <= 0 {
		c.RespuestaError("La cantidad debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()

	// Verificar stock
	var producto struct {
		Stock int `orm:"column(stock)"`
	}
	err := o.QueryTable("productos").Filter("id", carrito.ProductId).One(&producto)
	if err != nil {
		c.RespuestaError("Producto no encontrado", 404)
		return
	}
	if producto.Stock < carrito.Cantidad {
		c.RespuestaError("Stock insuficiente", 400)
		return
	}

	if _, err := o.Insert(&carrito); err != nil {
		c.RespuestaError("Error al crear carrito", 500)
		return
	}

	c.RespuestaExito(carrito)
}

func (c *ControladorCarrito) Listar() {
	var carritos []Carrito
	o := orm.NewOrm()
	qs := o.QueryTable("carritos")

	if userId, err := c.GetInt64("user_id"); err == nil {
		qs = qs.Filter("id_usuario", userId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&carritos); err != nil {
		c.RespuestaError("Error al obtener carritos", 500)
		return
	}

	c.RespuestaExito(carritos)
}

func (c *ControladorCarrito) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var carrito Carrito
	o := orm.NewOrm()
	if err := o.QueryTable("carritos").Filter("id_carrito", id).One(&carrito); err != nil {
		c.RespuestaError("Carrito no encontrado", 404)
		return
	}

	c.RespuestaExito(carrito)
}

func (c *ControladorCarrito) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var carrito Carrito
	if err := c.ParsearYValidarJSON(&carrito); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if carrito.Cantidad <= 0 {
		c.RespuestaError("La cantidad debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()

	// Verificar stock
	var producto struct {
		Stock int `orm:"column(stock)"`
	}
	err = o.QueryTable("productos").Filter("id", carrito.ProductId).One(&producto)
	if err != nil {
		c.RespuestaError("Producto no encontrado", 404)
		return
	}
	if producto.Stock < carrito.Cantidad {
		c.RespuestaError("Stock insuficiente", 400)
		return
	}

	carrito.Id = id
	if _, err := o.Update(&carrito); err != nil {
		c.RespuestaError("Error al actualizar carrito", 500)
		return
	}

	c.RespuestaExito(carrito)
}

func (c *ControladorCarrito) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Carrito{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar carrito", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Carrito eliminado"})
}
