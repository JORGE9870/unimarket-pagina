package repositorios

import models "unimarket/modelos"

type RepositorioMetricas interface {
	RecordProductProcess(sku string, status string)
	GetProductViews(productId string) int
	GetProductSales(productId string) int
	RecordStockMovement(movement *models.StockMovement)
}
