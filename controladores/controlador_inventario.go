package controladores

import (
	"unimarket/models"
	business "unimarket/servicios/negocio"

	"github.com/beego/beego/v2/client/orm"
)

type ControladorInventario struct {
	ControladorBase
	inventoryService *business.InventoryService
}

type MovimientoInventario struct {
	Id         int64  `json:"id_movimiento"`
	ProductId  int64  `json:"id_producto"`
	Cantidad   int    `json:"cantidad"`
	Tipo       string `json:"tipo"`
	CreateDate string `json:"fecha_creacion"`
}

// @router /movement [post]
func (c *ControladorInventario) ProcessMovement() {
	var movement models.StockMovement
	if err := c.ParsearYValidarJSON(&movement); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	// Verificar permisos de inventario
	if !c.HasInventoryPermission(movement.ProductID) {
		c.RespuestaError("Sin permiso para modificar inventario", 403)
		return
	}

	err := c.inventoryService.ProcessStockMovement(&movement)
	if err != nil {
		c.RespuestaError("Error al procesar movimiento", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Movimiento procesado correctamente"})
}

// @router /projection/:id [get]
func (c *ControladorInventario) GetProjection() {
	productId := c.Ctx.Input.Param(":id")
	days, err := c.GetInt("days", 30)
	if err != nil {
		c.RespuestaError("Días inválidos", 400)
		return
	}

	// Verificar permisos para ver proyecciones
	if !c.HasProjectionPermission() {
		c.RespuestaError("Sin acceso a proyecciones", 403)
		return
	}

	projection, err := c.inventoryService.GetInventoryProjection(productId, days)
	if err != nil {
		c.RespuestaError("Error al obtener proyección", 500)
		return
	}

	c.RespuestaExito(projection)
}

func (c *ControladorInventario) HasInventoryPermission(productId int64) bool {
	roles := c.Ctx.Input.GetData("roles").([]string)
	return containsAny(roles, []string{"admin", "inventory_manager"})
}

func (c *ControladorInventario) HasProjectionPermission() bool {
	roles := c.Ctx.Input.GetData("roles").([]string)
	return containsAny(roles, []string{"admin", "analyst", "inventory_manager"})
}

func containsAny(slice []string, targets []string) bool {
	for _, s := range slice {
		for _, t := range targets {
			if s == t {
				return true
			}
		}
	}
	return false
}

func (c *ControladorInventario) Crear() {
	var movimiento MovimientoInventario
	if err := c.ParsearYValidarJSON(&movimiento); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if movimiento.ProductId == 0 {
		c.RespuestaError("El producto es requerido", 400)
		return
	}
	if movimiento.Cantidad <= 0 {
		c.RespuestaError("La cantidad debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	// Actualizar stock
	_, err = tx.Raw("UPDATE productos SET stock = stock + ? WHERE id = ?",
		movimiento.Cantidad, movimiento.ProductId).Exec()
	if err != nil {
		tx.Rollback()
		c.RespuestaError("Error al actualizar stock", 500)
		return
	}

	if _, err := tx.Insert(&movimiento); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al registrar movimiento", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(movimiento)
}

func (c *ControladorInventario) Listar() {
	var movimientos []MovimientoInventario
	o := orm.NewOrm()
	qs := o.QueryTable("movimientos_inventario")

	if productId, err := c.GetInt64("product_id"); err == nil {
		qs = qs.Filter("id_producto", productId)
	}

	if _, err := qs.OrderBy("-fecha_creacion").All(&movimientos); err != nil {
		c.RespuestaError("Error al obtener movimientos", 500)
		return
	}

	c.RespuestaExito(movimientos)
}

func (c *ControladorInventario) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var movimiento MovimientoInventario
	o := orm.NewOrm()
	if err := o.QueryTable("movimientos_inventario").Filter("id_movimiento", id).One(&movimiento); err != nil {
		c.RespuestaError("Movimiento no encontrado", 404)
		return
	}

	c.RespuestaExito(movimiento)
}

func (c *ControladorInventario) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var movimiento MovimientoInventario
	if err := c.ParsearYValidarJSON(&movimiento); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if movimiento.Cantidad <= 0 {
		c.RespuestaError("La cantidad debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	// Obtener movimiento original
	var movimientoOriginal MovimientoInventario
	if err := tx.QueryTable("movimientos_inventario").Filter("id_movimiento", id).One(&movimientoOriginal); err != nil {
		tx.Rollback()
		c.RespuestaError("Movimiento no encontrado", 404)
		return
	}

	// Revertir stock anterior
	_, err = tx.Raw("UPDATE productos SET stock = stock - ? WHERE id = ?",
		movimientoOriginal.Cantidad, movimientoOriginal.ProductId).Exec()
	if err != nil {
		tx.Rollback()
		c.RespuestaError("Error al revertir stock", 500)
		return
	}

	// Aplicar nuevo stock
	_, err = tx.Raw("UPDATE productos SET stock = stock + ? WHERE id = ?",
		movimiento.Cantidad, movimiento.ProductId).Exec()
	if err != nil {
		tx.Rollback()
		c.RespuestaError("Error al actualizar stock", 500)
		return
	}

	movimiento.Id = id
	if _, err := tx.Update(&movimiento); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al actualizar movimiento", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(movimiento)
}

func (c *ControladorInventario) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		c.RespuestaError("Error al iniciar transacción", 500)
		return
	}

	// Obtener movimiento
	var movimiento MovimientoInventario
	if err := tx.QueryTable("movimientos_inventario").Filter("id_movimiento", id).One(&movimiento); err != nil {
		tx.Rollback()
		c.RespuestaError("Movimiento no encontrado", 404)
		return
	}

	// Revertir stock
	_, err = tx.Raw("UPDATE productos SET stock = stock - ? WHERE id = ?",
		movimiento.Cantidad, movimiento.ProductId).Exec()
	if err != nil {
		tx.Rollback()
		c.RespuestaError("Error al revertir stock", 500)
		return
	}

	if _, err := tx.Delete(&MovimientoInventario{Id: id}); err != nil {
		tx.Rollback()
		c.RespuestaError("Error al eliminar movimiento", 500)
		return
	}

	if err := tx.Commit(); err != nil {
		c.RespuestaError("Error al confirmar transacción", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Movimiento eliminado"})
}
