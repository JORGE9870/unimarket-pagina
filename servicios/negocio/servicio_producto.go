package business

import (
	"errors"
	"unimarket/models"
	"unimarket/repositories"
)

type ProductService struct {
	repo    repositories.ProductRepository
	cache   repositories.CacheRepository
	metrics repositories.MetricsRepository
}

func (s *ProductService) ProcessProduct(data *models.ProductRequest) (*models.ProductResponse, error) {
	// Validar reglas de negocio
	if err := s.validateBusinessRules(data); err != nil {
		return nil, err
	}

	// Procesar imágenes si existen
	if data.Images != nil {
		if err := s.processImages(data.Images); err != nil {
			return nil, err
		}
	}

	// Calcular precios y descuentos
	finalPrice := s.calculateFinalPrice(data.BasePrice, data.Discount)

	// Verificar stock en múltiples ubicaciones
	if err := s.verifyStockAvailability(data.SKU, data.Quantity); err != nil {
		return nil, err
	}

	// Crear respuesta procesada
	response := &models.ProductResponse{
		Name:       data.Name,
		FinalPrice: finalPrice,
		Available:  true,
		Locations:  s.getAvailableLocations(data.SKU),
	}

	// Registrar métricas de procesamiento
	s.metrics.RecordProductProcess(data.SKU, "success")

	return response, nil
}

func (s *ProductService) validateBusinessRules(data *models.ProductRequest) error {
	// Validar categorías permitidas
	if !s.areValidCategories(data.Categories) {
		return errors.New("categorías no válidas para el tipo de producto")
	}

	// Validar restricciones de precio
	if !s.isPriceValid(data.BasePrice, data.Categories) {
		return errors.New("precio fuera del rango permitido para las categorías")
	}

	// Validar restricciones de marca
	if !s.isBrandAllowed(data.Brand, data.Categories) {
		return errors.New("marca no autorizada para estas categorías")
	}

	return nil
}

func (s *ProductService) GetProductAnalytics(productId string) (*models.ProductAnalytics, error) {
	// Intentar obtener del caché
	analytics, err := s.cache.GetProductAnalytics(productId)
	if err == nil {
		return analytics, nil
	}

	// Calcular analíticas
	analytics = &models.ProductAnalytics{
		ViewCount:     s.metrics.GetProductViews(productId),
		SalesCount:    s.metrics.GetProductSales(productId),
		ReviewsAvg:    s.calculateReviewsAverage(productId),
		StockTurnover: s.calculateStockTurnover(productId),
		TrendScore:    s.calculateTrendScore(productId),
	}

	// Guardar en caché
	s.cache.SetProductAnalytics(productId, analytics)

	return analytics, nil
}

func (s *ProductService) UpdateProductStatus(productId string, status string) error {
	// Validar transición de estado
	if !s.isValidStatusTransition(productId, status) {
		return errors.New("transición de estado no permitida")
	}

	// Actualizar estado
	if err := s.repo.UpdateStatus(productId, status); err != nil {
		return err
	}

	// Notificar a sistemas externos
	s.notifyExternalSystems(productId, status)

	return nil
}
