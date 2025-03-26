package services

import (
	"strconv"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type NotificationService struct {
	orm orm.Ormer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		orm: orm.NewOrm(),
	}
}

func (s *NotificationService) CreateOrderNotification(userId int64, orderId int64, status string) error {
	notification := &models.Notification{
		UserId:  userId,
		Title:   "Actualización de Pedido",
		Message: "Tu pedido #" + strconv.FormatInt(orderId, 10) + " ha cambiado a estado: " + status,
		Type:    "order_status",
	}

	_, err := s.orm.Insert(notification)
	return err
}

func (s *NotificationService) CreateDeliveryNotification(userId int64, orderId int64, status string) error {
	notification := &models.Notification{
		UserId:  userId,
		Title:   "Actualización de Entrega",
		Message: "La entrega de tu pedido #" + strconv.FormatInt(orderId, 10) + " está " + status,
		Type:    "delivery_status",
	}

	_, err := s.orm.Insert(notification)
	return err
}

func (s *NotificationService) GetUnreadCount(userId int64) (int64, error) {
	return s.orm.QueryTable("notificaciones").
		Filter("id_usuario", userId).
		Filter("leida", false).
		Count()
}
