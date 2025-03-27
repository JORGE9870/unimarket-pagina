package repositorios

import "unimarket/models"

type RepositorioInventario interface {
	RecordMovement(movement *models.StockMovement) error
	GetHistoricalMovements(productId string) []models.StockMovement
}
