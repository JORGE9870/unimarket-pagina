package controllers

import (
	"encoding/json"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type OrderController struct {
	BaseController
}

// @Title CreateOrder
// @Description create new order
// @Success 200 {object} models.Order
// @router / [post]
func (o *OrderController) Create() {
	var order models.Order
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &order); err != nil {
		o.ResponseError("Datos de solicitud inválidos", 400)
		return
	}

	// Validaciones básicas
	if order.User == nil || order.User.Id == 0 {
		o.ResponseError("El usuario es requerido", 400)
		return
	}
	if order.Store == nil || order.Store.Id == 0 {
		o.ResponseError("La tienda es requerida", 400)
		return
	}
	if len(order.Items) == 0 {
		o.ResponseError("El pedido debe contener al menos un producto", 400)
		return
	}

	oo := orm.NewOrm()

	// Validar stock y calcular total
	var total float64
	for _, item := range order.Items {
		if item.Product == nil || item.Product.Id == 0 {
			o.ResponseError("Producto inválido", 400)
			return
		}

		// Cargar producto para validar stock
		var product models.Product
		if err := oo.Read(item.Product); err != nil {
			o.ResponseError("Producto no encontrado", 404)
			return
		}

		if product.Stock < item.Quantity {
			o.ResponseError("Stock insuficiente para el producto: "+product.Name, 400)
			return
		}

		// Actualizar stock
		product.Stock -= item.Quantity
		if _, err := oo.Update(&product); err != nil {
			o.ResponseError("Error al actualizar el stock", 500)
			return
		}

		// Establecer precios
		item.UnitPrice = product.Price
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		total += item.TotalPrice
	}

	// Establecer total y estado inicial
	order.TotalAmount = total
	order.Status = "pending"

	// Iniciar transacción
	tx, err := oo.Begin()
	if err != nil {
		o.ResponseError("Error al iniciar la transacción", 500)
		return
	}

	// Insertar el pedido
	if _, err := tx.Insert(&order); err != nil {
		tx.Rollback()
		o.ResponseError("Error al crear el pedido", 500)
		return
	}

	// Insertar los items del pedido
	for _, item := range order.Items {
		item.Order = &order
		if _, err := tx.Insert(item); err != nil {
			tx.Rollback()
			o.ResponseError("Error al crear los items del pedido", 500)
			return
		}
	}

	// Crear la entrega
	if order.Delivery != nil {
		order.Delivery.Order = &order
		if _, err := tx.Insert(order.Delivery); err != nil {
			tx.Rollback()
			o.ResponseError("Error al crear la entrega", 500)
			return
		}
	}

	// Confirmar transacción
	if err := tx.Commit(); err != nil {
		o.ResponseError("Error al confirmar la transacción", 500)
		return
	}

	o.ResponseSuccess(order)
}

// @Title GetAllOrders
// @Description get all orders with optional filters
// @Success 200 {object} []models.Order
// @router / [get]
func (o *OrderController) GetAll() {
	var orders []models.Order
	oo := orm.NewOrm()
	qs := oo.QueryTable(new(models.Order))

	// Aplicar filtros si se proporcionan
	if userId, err := o.GetInt64("user_id"); err == nil {
		qs = qs.Filter("User__Id", userId)
	}
	if storeId, err := o.GetInt64("store_id"); err == nil {
		qs = qs.Filter("Store__Id", storeId)
	}
	if status := o.GetString("status"); status != "" {
		qs = qs.Filter("Status", status)
	}

	if _, err := qs.All(&orders); err != nil {
		o.ResponseError("Error al obtener los pedidos", 500)
		return
	}

	// Cargar relaciones
	for i := range orders {
		oo.LoadRelated(&orders[i], "Items")
		oo.LoadRelated(&orders[i], "Delivery")
	}

	o.ResponseSuccess(orders)
}

// @Title GetOrder
// @Description get order by id
// @Success 200 {object} models.Order
// @router /:id [get]
func (o *OrderController) Get() {
	id, err := o.GetInt64(":id")
	if err != nil {
		o.ResponseError("ID inválido", 400)
		return
	}

	var order models.Order
	oo := orm.NewOrm()
	if err := oo.QueryTable(new(models.Order)).Filter("Id", id).One(&order); err != nil {
		o.ResponseError("Pedido no encontrado", 404)
		return
	}

	// Cargar relaciones
	oo.LoadRelated(&order, "Items")
	oo.LoadRelated(&order, "Delivery")

	o.ResponseSuccess(order)
}

// @Title UpdateOrder
// @Description update order
// @Success 200 {object} models.Order
// @router /:id [put]
func (o *OrderController) Update() {
	id, err := o.GetInt64(":id")
	if err != nil {
		o.ResponseError("ID inválido", 400)
		return
	}

	var order models.Order
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &order); err != nil {
		o.ResponseError("Datos de solicitud inválidos", 400)
		return
	}

	// Validaciones básicas
	if order.Status == "" {
		o.ResponseError("El estado del pedido es requerido", 400)
		return
	}

	order.Id = id
	oo := orm.NewOrm()

	// Iniciar transacción
	tx, err := oo.Begin()
	if err != nil {
		o.ResponseError("Error al iniciar la transacción", 500)
		return
	}

	// Actualizar el pedido
	if _, err := tx.Update(&order); err != nil {
		tx.Rollback()
		o.ResponseError("Error al actualizar el pedido", 500)
		return
	}

	// Si hay items nuevos, actualizarlos
	if len(order.Items) > 0 {
		// Eliminar items antiguos
		if _, err := tx.QueryTable(new(models.OrderItem)).Filter("Order__Id", id).Delete(); err != nil {
			tx.Rollback()
			o.ResponseError("Error al eliminar los items antiguos", 500)
			return
		}

		// Insertar nuevos items
		for _, item := range order.Items {
			item.Order = &order
			if _, err := tx.Insert(item); err != nil {
				tx.Rollback()
				o.ResponseError("Error al actualizar los items del pedido", 500)
				return
			}
		}
	}

	// Si hay datos de entrega, actualizarlos
	if order.Delivery != nil {
		order.Delivery.Order = &order
		if _, err := tx.Update(order.Delivery); err != nil {
			tx.Rollback()
			o.ResponseError("Error al actualizar la entrega", 500)
			return
		}
	}

	// Confirmar transacción
	if err := tx.Commit(); err != nil {
		o.ResponseError("Error al confirmar la transacción", 500)
		return
	}

	o.ResponseSuccess(order)
}

// @Title DeleteOrder
// @Description delete order
// @Success 200 {object} models.Order
// @router /:id [delete]
func (o *OrderController) Delete() {
	id, err := o.GetInt64(":id")
	if err != nil {
		o.ResponseError("ID inválido", 400)
		return
	}

	oo := orm.NewOrm()

	// Iniciar transacción
	tx, err := oo.Begin()
	if err != nil {
		o.ResponseError("Error al iniciar la transacción", 500)
		return
	}

	// Eliminar items del pedido
	if _, err := tx.QueryTable(new(models.OrderItem)).Filter("Order__Id", id).Delete(); err != nil {
		tx.Rollback()
		o.ResponseError("Error al eliminar los items del pedido", 500)
		return
	}

	// Eliminar datos de entrega
	if _, err := tx.QueryTable(new(models.Delivery)).Filter("Order__Id", id).Delete(); err != nil {
		tx.Rollback()
		o.ResponseError("Error al eliminar los datos de entrega", 500)
		return
	}

	// Eliminar el pedido
	if _, err := tx.Delete(&models.Order{BaseModel: models.BaseModel{Id: id}}); err != nil {
		tx.Rollback()
		o.ResponseError("Error al eliminar el pedido", 500)
		return
	}

	// Confirmar transacción
	if err := tx.Commit(); err != nil {
		o.ResponseError("Error al confirmar la transacción", 500)
		return
	}

	o.ResponseSuccess(map[string]string{"message": "Pedido eliminado exitosamente"})
}
