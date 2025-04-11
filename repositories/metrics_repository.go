package repositories

type MetricsRepository interface {
	RecordProductProcess(sku string, status string)
	GetProductViews(productId string) int
	GetProductSales(productId string) int
}
