package controllers

import (
	"encoding/json"
	"unimarket/models"
)

type StoreController struct {
	BaseController
}

// @Title CreateStore
// @Description create new store
// @Success 200 {object} models.Store
// @router / [post]
func (s *StoreController) Create() {
	var store models.Store
	if err := json.Unmarshal(s.Ctx.Input.RequestBody, &store); err != nil {
		s.ResponseError("Datos de solicitud inválidos", 400)
		return
	}
	s.ResponseSuccess(store)
}

// @Title GetAllStores
// @Description get all stores
// @Success 200 {object} []models.Store
// @router / [get]
func (s *StoreController) GetAll() {
	s.ResponseSuccess([]models.Store{})
}

// @Title GetStore
// @Description get store by id
// @Success 200 {object} models.Store
// @router /:id [get]
func (s *StoreController) Get() {
	id := s.GetString(":id")
	s.ResponseSuccess(map[string]string{"id": id})
}

// @Title UpdateStore
// @Description update store
// @Success 200 {object} models.Store
// @router /:id [put]
func (s *StoreController) Update() {
	id := s.GetString(":id")
	s.ResponseSuccess(map[string]string{"id": id, "message": "Store updated"})
}

// @Title DeleteStore
// @Description delete store
// @Success 200 {object} models.Store
// @router /:id [delete]
func (s *StoreController) Delete() {
	id := s.GetString(":id")
	s.ResponseSuccess(map[string]string{"id": id, "message": "Store deleted"})
}
