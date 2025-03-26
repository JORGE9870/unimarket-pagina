package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

type InventoryController struct {
	BaseController
}

type InventoryMovement struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id"`
	Type      string `json:"type"` // entrada/salida
	Quantity  int    `json:"quantity"`
	Reason    string `json:"reason"`
	Date      string `json:"date"`
}

func (i *InventoryController) RegisterMovement() {
	var movement InventoryMovement
	if err := json.Unmarshal(i.Ctx.Input.RequestBody, &movement); err != nil {
		i.ResponseError("Datos inválidos", 400)
		return
	}

	if movement.Quantity <= 0 {
		i.ResponseError("La cantidad debe ser mayor a 0", 400)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&movement); err != nil {
		i.ResponseError("Error al registrar movimiento", 500)
		return
	}

	i.ResponseSuccess(movement)
}
