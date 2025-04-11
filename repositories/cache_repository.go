package repositories

import models "unimarket/modelos"

type CacheRepository interface {
	GetProductAnalytics(productId string) (*models.ProductAnalytics, error)
	SetProductAnalytics(productId string, analytics *models.ProductAnalytics) error
}
