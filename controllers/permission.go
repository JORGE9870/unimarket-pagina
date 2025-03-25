package controllers

import (
	"encoding/json"
	"unimarket/models"
)

type PermissionController struct {
	BaseController
}

// @Title CreatePermission
// @Description create new permission
// @Success 200 {object} models.Permission
// @router / [post]
func (p *PermissionController) Create() {
	var permission models.Permission
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &permission); err != nil {
		p.ResponseError("Datos de solicitud inválidos", 400)
		return
	}
	p.ResponseSuccess(permission)
}

// @Title GetAllPermissions
// @Description get all permissions
// @Success 200 {object} []models.Permission
// @router / [get]
func (p *PermissionController) GetAll() {
	p.ResponseSuccess([]models.Permission{})
}

// @Title GetPermission
// @Description get permission by id
// @Success 200 {object} models.Permission
// @router /:id [get]
func (p *PermissionController) Get() {
	id := p.GetString(":id")
	p.ResponseSuccess(map[string]string{"id": id})
}

// @Title UpdatePermission
// @Description update permission
// @Success 200 {object} models.Permission
// @router /:id [put]
func (p *PermissionController) Update() {
	id := p.GetString(":id")
	p.ResponseSuccess(map[string]string{"id": id, "message": "Permission updated"})
}

// @Title DeletePermission
// @Description delete permission
// @Success 200 {object} models.Permission
// @router /:id [delete]
func (p *PermissionController) Delete() {
	id := p.GetString(":id")
	p.ResponseSuccess(map[string]string{"id": id, "message": "Permission deleted"})
}
