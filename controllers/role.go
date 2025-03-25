package controllers

import (
	"encoding/json"
	"unimarket/models"
)

type RoleController struct {
	BaseController
}

// @Title CreateRole
// @Description create new role
// @Success 200 {object} models.Role
// @router / [post]
func (r *RoleController) Create() {
	var role models.Role
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &role); err != nil {
		r.ResponseError("Datos de solicitud inválidos", 400)
		return
	}
	r.ResponseSuccess(role)
}

// @Title GetAllRoles
// @Description get all roles
// @Success 200 {object} []models.Role
// @router / [get]
func (r *RoleController) GetAll() {
	r.ResponseSuccess([]models.Role{})
}

// @Title GetRole
// @Description get role by id
// @Success 200 {object} models.Role
// @router /:id [get]
func (r *RoleController) Get() {
	id := r.GetString(":id")
	r.ResponseSuccess(map[string]string{"id": id})
}

// @Title UpdateRole
// @Description update role
// @Success 200 {object} models.Role
// @router /:id [put]
func (r *RoleController) Update() {
	id := r.GetString(":id")
	r.ResponseSuccess(map[string]string{"id": id, "message": "Role updated"})
}

// @Title DeleteRole
// @Description delete role
// @Success 200 {object} models.Role
// @router /:id [delete]
func (r *RoleController) Delete() {
	id := r.GetString(":id")
	r.ResponseSuccess(map[string]string{"id": id, "message": "Role deleted"})
}
