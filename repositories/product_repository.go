package repositories

type ProductRepository interface {
	UpdateStatus(productId string, status string) error
}
