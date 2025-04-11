package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorPermiso struct {
	ControladorBase
}

type Permiso struct {
	Id          int64  `json:"id_permiso"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	CreateDate  string `json:"fecha_creacion"`
}

// @Title CreatePermission
// @Description create new permission
// @Success 200 {object} models.Permission
// @router / [post]
func (c *ControladorPermiso) Crear() {
	var permiso Permiso
	if err := c.ParsearYValidarJSON(&permiso); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}
	func (c *ControladorPermiso) Crear() {
		var permiso Permiso
		if err := c.ParsearYValidarJSON(&permiso); err != nil {
			c.RespuestaError("Datos inválidos", 400)
			return
		}
	
	if permiso.Nombre == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&permiso); err != nil {
		c.RespuestaError("Error al crear permiso", 500)
		return
	}

	c.RespuestaExito(permiso)
}

// @Title GetAllPermissions
// @Description get all permissions
// @Success 200 {object} []models.Permission
// @router / [get]
func (c *ControladorPermiso) Listar() {
	var permisos []Permiso
	o := orm.NewOrm()

	if _, err := o.QueryTable("permisos").OrderBy("-fecha_creacion").All(&permisos); err != nil {
		c.RespuestaError("Error al obtener permisos", 500)
		return
	}

	c.RespuestaExito(permisos)
}

// @Title GetPermission
// @Description get permission by id
// @Success 200 {object} models.Permission
// @router /:id [get]
func (c *ControladorPermiso) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var permiso Permiso
	o := orm.NewOrm()
	if err := o.QueryTable("permisos").Filter("id_permiso", id).One(&permiso); err != nil {
		c.RespuestaError("Permiso no encontrado", 404)
		return
	}

	c.RespuestaExito(permiso)
}

// @Title UpdatePermission
// @Description update permission
// @Success 200 {object} models.Permission
// @router /:id [put]
func (c *ControladorPermiso) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var permiso Permiso
	if err := c.ParsearYValidarJSON(&permiso); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if permiso.Nombre == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}

	permiso.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&permiso); err != nil {
		c.RespuestaError("Error al actualizar permiso", 500)
		return
	}

	c.RespuestaExito(permiso)
}

// @Title DeletePermission
// @Description delete permission
// @Success 200 {object} models.Permission
// @router /:id [delete]
func (c *ControladorPermiso) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Permiso{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar permiso", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Permiso eliminado"})
}
