package repositorios

type RepositorioNotificaciones interface {
	SendLowStockAlert(productId int64)
}
