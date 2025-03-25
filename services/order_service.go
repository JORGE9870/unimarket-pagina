package services

import (
	"errors"
	"unimarket/models"

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

func (s *OrderService) Create(order *models.Order) error {
	tx, err := s.orm.Begin()
	if err != nil {
		return err
	}

	// Validar stock de productos
	for _, item := range order.Items {
		var product models.Product
		err := tx.QueryTable("productos").Filter("id_producto", item.Product.Id).One(&product)
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
		item.Order = order
		if _, err := tx.Insert(item); err != nil {
			tx.Rollback()
			return err
		}

		// Actualizar stock
		if _, err := tx.QueryTable("productos").Filter("id_producto", item.Product.Id).Update(orm.Params{
			"stock": orm.ColValue(orm.ColMinus, item.Quantity),
		}); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
