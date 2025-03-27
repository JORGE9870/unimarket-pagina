package repositorios

import "unimarket/models"

type RepositorioMetricas interface {
	RecordStockMovement(movement *models.StockMovement)
}
