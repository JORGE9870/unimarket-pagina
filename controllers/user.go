package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type UserController struct {
	BaseController
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
func (u *UserController) Post() {
	var user User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		u.ResponseError("Datos de solicitud inválidos", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&user); err != nil {
		u.ResponseError("Error al crear usuario", 500)
		return
	}

	u.ResponseSuccess(user)
}

// @Title GetUser
// @Description get user by id
// @Param   id        path    int     true        "user id"
// @Success 200 {object} models.User
// @Failure 404 User not found
// @router /:id [get]
func (u *UserController) Get() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.ResponseError("ID inválido", 400)
		return
	}

	var user User
	o := orm.NewOrm()
	if err := o.QueryTable("usuarios").Filter("id_usuario", id).One(&user); err != nil {
		u.ResponseError("Usuario no encontrado", 404)
		return
	}

	u.ResponseSuccess(user)
}

func (u *UserController) GetAll() {
	var users []User
	o := orm.NewOrm()

	if _, err := o.QueryTable("usuarios").All(&users); err != nil {
		u.ResponseError("Error al obtener usuarios", 500)
		return
	}

	u.ResponseSuccess(users)
}

func (u *UserController) Update() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.ResponseError("ID inválido", 400)
		return
	}

	var user User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		u.ResponseError("Datos inválidos", 400)
		return
	}

	user.Id = id
	o := orm.NewOrm()
	if _, err := o.Update(&user); err != nil {
		u.ResponseError("Error al actualizar usuario", 500)
		return
	}

	u.ResponseSuccess(user)
}

func (u *UserController) Delete() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.ResponseError("ID inválido", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&User{Id: id}); err != nil {
		u.ResponseError("Error al eliminar usuario", 500)
		return
	}

	u.ResponseSuccess(map[string]string{"message": "Usuario eliminado"})
}
