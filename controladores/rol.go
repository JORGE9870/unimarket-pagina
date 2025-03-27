package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorRol struct {
	ControladorBase
}

type Rol struct {
	Id          int64  `json:"id_rol"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	CreateDate  string `json:"fecha_creacion"`
}

// @Title CreateRole
// @Description create new role
// @Success 200 {object} models.Role
// @router / [post]
func (c *ControladorRol) Crear() {
	var rol Rol
	if err := c.ParsearYValidarJSON(&rol); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if rol.Nombre == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&rol); err != nil {
		c.RespuestaError("Error al crear rol", 500)
		return
	}

	c.RespuestaExito(rol)
}

// @Title GetAllRoles
// @Description get all roles
// @Success 200 {object} []models.Role
// @router / [get]
func (c *ControladorRol) Listar() {
	var roles []Rol
	o := orm.NewOrm()

	if _, err := o.QueryTable("roles").OrderBy("-fecha_creacion").All(&roles); err != nil {
		c.RespuestaError("Error al obtener roles", 500)
		return
	}

	c.RespuestaExito(roles)
}

// @Title GetRole
// @Description get role by id
// @Success 200 {object} models.Role
// @router /:id [get]
func (c *ControladorRol) Obtener() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var rol Rol
	o := orm.NewOrm()
	if err := o.QueryTable("roles").Filter("id_rol", id).One(&rol); err != nil {
		c.RespuestaError("Rol no encontrado", 404)
		return
	}

	c.RespuestaExito(rol)
}

// @Title UpdateRole
// @Description update role
// @Success 200 {object} models.Role
// @router /:id [put]
func (c *ControladorRol) Actualizar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	var rol Rol
	if err := c.ParsearYValidarJSON(&rol); err != nil {
		c.RespuestaError("Datos inválidos", 400)
		return
	}

	if rol.Nombre == "" {
		c.RespuestaError("El nombre es requerido", 400)
		return
	}

	rol.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&rol); err != nil {
		c.RespuestaError("Error al actualizar rol", 500)
		return
	}

	c.RespuestaExito(rol)
}

// @Title DeleteRole
// @Description delete role
// @Success 200 {object} models.Role
// @router /:id [delete]
func (c *ControladorRol) Eliminar() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&Rol{Id: id}); err != nil {
		c.RespuestaError("Error al eliminar rol", 500)
		return
	}

	c.RespuestaExito(map[string]string{"mensaje": "Rol eliminado"})
}
