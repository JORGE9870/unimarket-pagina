package controladores

import (
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type Controlador_tienda struct {
	ControladorBase
}

type tienda struct {
	Id          int64  `json:"id_tienda"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Direccion   string `json:"direccion"`
	CreateDate  string `json:"fecha_creacion"`
}

// @Title CreateStore
// @Description create new store
// @Success 200 {object} models.Store
// @router / [post]
func (c *ControladorTienda) Crear() {
	var tienda Tienda
	if err := c.ParsearYValidarJSON(&tienda); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if tienda.Nombre == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&tienda); err != nil {
		c.RespuestaError("Error al crear tienda", 500)
		return
	}

	c.RespuestaExito(tienda)
}

// @Title GetAllStores
// @Description get all stores
// @Success 200 {object} []models.Tienda
// @router / [get]
func (c *ControladorTienda) Listar() {
	c.RespuestaExito([]models.Tienda{})
}

// @Title GetStore
// @Description get store by id
// @Success 200 {object} models.Store
// @router /:id [get]
func (c *ControladorTienda) Obtener() {
	id := c.GetString(":id")
	c.RespuestaExito(map[string]string{"id": id})
}

// @Title UpdateStore
// @Description update store
// @Success 200 {object} models.Store
// @router /:id [put]
func (c *ControladorTienda) Actualizar() {
	id := c.GetString(":id")
	c.RespuestaExito(map[string]string{"id": id, "mensaje": "Tienda actualizada"})
}

func (c *ControladorTienda) Eliminar() {
	id := c.GetString(":id")
	c.RespuestaExito(map[string]string{"id": id, "mensaje": "Tienda eliminada"})
}
