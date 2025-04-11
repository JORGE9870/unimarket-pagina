package business

import (
	"errors"
	models "unimarket/modelos"
	"unimarket/servicios/repositorios"
)

type InventoryService struct {
	repo     repositorios.RepositorioInventario
	metrics  repositorios.RepositorioMetricas
	notifier repositorios.RepositorioNotificaciones
}

func (s *InventoryService) validateMovement(movement *models.StockMovement) error {
	if movement == nil {
		return errors.New("movimiento no puede ser nulo")
	}
	if movement.ProductID <= 0 {
		return errors.New("id de producto inválido")
	}
	if movement.Quantity <= 0 {
		return errors.New("cantidad debe ser mayor a 0")
	}
	return nil
}

func (s *InventoryService) shouldTriggerAlert(movement *models.StockMovement) bool {
	return movement.Quantity < 10 // Ejemplo simple de alerta de bajo stock
}

func (s *InventoryService) calculateInventoryTrend(movements []models.StockMovement) float64 {
	if len(movements) == 0 {
		return 0
	}
	// Implementación simple de cálculo de tendencia
	return 0.0
}

func (s *InventoryService) calculateFutureStock(trend float64, days int) *models.InventoryProjection {
	return &models.InventoryProjection{
		Days: days,
		// Otros campos según necesidad
	}
}

func (s *InventoryService) ProcessStockMovement(movement *models.StockMovement) error {
	// Validar movimiento
	if err := s.validateMovement(movement); err != nil {
		return err
	}

	// Verificar límites y alertas
	if s.shouldTriggerAlert(movement) {
		s.notifier.SendLowStockAlert(movement.ProductID)
	}

	// Registrar movimiento
	if err := s.repo.RecordMovement(movement); err != nil {
		return err
	}

	// Actualizar métricas
	s.metrics.RecordStockMovement(movement)

	return nil
}

func (s *InventoryService) GetInventoryProjection(productId string, days int) (*models.InventoryProjection, error) {
	// Obtener histórico de movimientos
	movements := s.repo.GetHistoricalMovements(productId)

	// Calcular tendencias
	trend := s.calculateInventoryTrend(movements)

	// Proyectar stock futuro
	projection := s.calculateFutureStock(trend, days)

	return projection, nil
}
