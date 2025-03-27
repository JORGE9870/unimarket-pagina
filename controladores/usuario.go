package controladores

import (
	"github.com/beego/beego/v2/client/orm"
)

type ControladorUsuario struct {
	ControladorBase
}

type User struct {
	Id           int64  `json:"id_usuario"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"nombre"`
	LastName     string `json:"apellido"`
	Phone        string `json:"telefono"`
	RegisterDate string `json:"fecha_registro"`
	Status       string `json:"estado"`
	Roles        []Role `json:"roles,omitempty"`
}

type Role struct {
	Id          int64  `json:"id_rol"`
	Name        string `json:"nombre"`
	Description string `json:"descripcion"`
}

// @Title CreateUser
// @Description create new user
// @Param   body        body    models.User   true        "user info"
// @Success 200 {object} models.User
// @Failure 400 Bad Request
// @router / [post]
func (u *ControladorUsuario) Crear() {
	var usuario User
	if err := u.ParsearYValidarJSON(&usuario); err != nil {
		u.RespuestaError("Datos inválidos", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&usuario); err != nil {
		u.RespuestaError("Error al crear usuario", 500)
		return
	}

	u.RespuestaExito(usuario)
}

// @Title GetUser
// @Description get user by id
// @Param   id        path    int     true        "user id"
// @Success 200 {object} models.User
// @Failure 404 User not found
// @router /:id [get]
func (u *ControladorUsuario) Obtener() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.RespuestaError("ID inválido", 400)
		return
	}

	var usuario User
	o := orm.NewOrm()
	if err := o.QueryTable("usuarios").Filter("id_usuario", id).One(&usuario); err != nil {
		u.RespuestaError("Usuario no encontrado", 404)
		return
	}

	u.RespuestaExito(usuario)
}

func (u *ControladorUsuario) Listar() {
	var usuarios []User
	o := orm.NewOrm()
	if _, err := o.QueryTable("usuarios").All(&usuarios); err != nil {
		u.RespuestaError("Error al obtener usuarios", 500)
		return
	}

	u.RespuestaExito(usuarios)
}

func (u *ControladorUsuario) Actualizar() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.RespuestaError("ID inválido", 400)
		return
	}

	var usuario User
	if err := u.ParsearYValidarJSON(&usuario); err != nil {
		u.RespuestaError("Datos inválidos", 400)
		return
	}

	usuario.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&usuario); err != nil {
		u.RespuestaError("Error al actualizar usuario", 500)
		return
	}

	u.RespuestaExito(usuario)
}

func (u *ControladorUsuario) Eliminar() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.RespuestaError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&User{Id: id}); err != nil {
		u.RespuestaError("Error al eliminar usuario", 500)
		return
	}

	u.RespuestaExito(map[string]string{"mensaje": "Usuario eliminado"})
}
