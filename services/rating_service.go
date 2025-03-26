package services

import (
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type RatingService struct {
	orm orm.Ormer
}

func NewRatingService() *RatingService {
	return &RatingService{
		orm: orm.NewOrm(),
	}
}

func (s *RatingService) GetProductAverageRating(productId int64) (float64, error) {
	var avg float64
	err := s.orm.Raw("SELECT AVG(puntuacion) FROM valoraciones WHERE id_producto = ?", productId).QueryRow(&avg)
	return avg, err
}

func (s *RatingService) HasUserRated(userId, productId int64) (bool, error) {
	exist := s.orm.QueryTable("valoraciones").
		Filter("id_usuario", userId).
		Filter("id_producto", productId).
		Exist()
	return exist, nil
}

func (s *RatingService) GetTopRatedProducts(limit int) ([]models.Product, error) {
	var products []models.Product
	_, err := s.orm.Raw(`
        SELECT p.*, AVG(r.puntuacion) as rating_promedio 
        FROM productos p 
        LEFT JOIN valoraciones r ON p.id_producto = r.id_producto 
        GROUP BY p.id_producto 
        ORDER BY rating_promedio DESC 
        LIMIT ?
    `, limit).QueryRows(&products)

	return products, err
}
