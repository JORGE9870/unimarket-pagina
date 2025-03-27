package services

import (
	"errors"
	"unimarket/controllers"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type CartService struct {
	orm orm.Ormer
}

func NewCartService() *CartService {
	return &CartService{
		orm: orm.NewOrm(),
	}
}

func (s *CartService) AddItem(cartId int64, item *controllers.CartItem) error {
	// Validar stock disponible
	var product models.Product
	err := s.orm.QueryTable("productos").Filter("id_producto", item.ProductId).One(&product)
	if err != nil {
		return errors.New("producto no encontrado")
	}
	if product.Stock < item.Quantity {
		return errors.New("stock insuficiente")
	}

	// Actualizar si ya existe el item
	var existingItem controllers.CartItem
	err = s.orm.QueryTable("items_carrito").
		Filter("id_carrito", cartId).
		Filter("id_producto", item.ProductId).
		One(&existingItem)

	if err == nil {
		// Actualizar cantidad
		existingItem.Quantity += item.Quantity
		if existingItem.Quantity > product.Stock {
			return errors.New("stock insuficiente")
		}
		_, err = s.orm.Update(&existingItem)
		return err
	}

	// Crear nuevo item
	item.CartId = cartId
	_, err = s.orm.Insert(item)
	return err
}
