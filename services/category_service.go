package services

import (
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type CategoryService struct {
	orm orm.Ormer
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		orm: orm.NewOrm(),
	}
}

func (s *CategoryService) GetCategoryTree() ([]models.Category, error) {
	var categories []models.Category
	_, err := s.orm.QueryTable("categorias").
		Filter("id_categoria_padre__isnull", true).
		All(&categories)

	if err != nil {
		return nil, err
	}

	for i := range categories {
		if err := s.loadSubcategories(&categories[i]); err != nil {
			return nil, err
		}
	}

	return categories, nil
}

func (s *CategoryService) loadSubcategories(category *models.Category) error {
	var subcategories []*models.Category
	_, err := s.orm.QueryTable("categorias").
		Filter("id_categoria_padre", category.Id).
		All(&subcategories)

	if err != nil {
		return err
	}

	category.Subcategories = subcategories
	for _, subcat := range subcategories {
		if err := s.loadSubcategories(subcat); err != nil {
			return err
		}
	}

	return nil
}

func (s *CategoryService) GetCategoryProducts(categoryId int64, page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product
	qs := s.orm.QueryTable("productos").
		Filter("Categories__CategoryId", categoryId)

	total, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = qs.Limit(pageSize).Offset((page - 1) * pageSize).All(&products)
	return products, total, err
}
