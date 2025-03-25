package services

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type OrderService struct {
	orm orm.Ormer
}

func NewOrderService() *OrderService {
	return &OrderService{
		orm: orm.NewOrm(),
	}
}

func (s *OrderService) Create(order *Order) error {
	tx, err := s.orm.Begin()
	if err != nil {
		return err
	}

	// Validar stock de productos
	for _, item := range order.Items {
		var product Product
		err := tx.QueryTable("productos").Filter("id_producto", item.ProductId).One(&product)
		if err != nil {
			tx.Rollback()
			return errors.New("producto no encontrado")
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			return errors.New("stock insuficiente")
		}
	}

	// Crear pedido
	if _, err := tx.Insert(order); err != nil {
		tx.Rollback()
		return err
	}

	// Crear items y actualizar stock
	for _, item := range order.Items {
		item.OrderId = order.Id
		if _, err := tx.Insert(item); err != nil {
			tx.Rollback()
			return err
		}

		// Actualizar stock
		if _, err := tx.QueryTable("productos").Filter("id_producto", item.ProductId).Update(orm.Params{
			"stock": orm.ColValue(orm.ColMinus, item.Quantity),
		}); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
