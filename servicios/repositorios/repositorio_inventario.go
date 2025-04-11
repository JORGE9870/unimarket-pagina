package repositorios

import "unimarket/modelos"

type RepositorioInventario interface {
	RecordMovement(movement *modelos.StockMovement) error
	GetHistoricalMovements(productId string) []modelos.StockMovement
}
