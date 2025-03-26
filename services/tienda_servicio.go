package services

import (
	"errors"
	"unimarket/models"

	"github.com/beego/beego/v2/client/orm"
)

type StoreService struct {
	orm orm.Ormer
}

func NewStoreService() *StoreService {
	return &StoreService{
		orm: orm.NewOrm(),
	}
}

func (s *StoreService) Create(store *models.Store) error {
	// Validar nombre único
	exist := s.orm.QueryTable("tiendas").Filter("nombre", store.Name).Exist()
	if exist {
		return errors.New("ya existe una tienda con ese nombre")
	}

	_, err := s.orm.Insert(store)
	return err
}

func (s *StoreService) GetStoreProducts(storeId int64, page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	qs := s.orm.QueryTable("productos").Filter("id_negocio", storeId)

	total, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = qs.Limit(pageSize).Offset((page - 1) * pageSize).All(&products)
	return products, total, err
}

func (s *StoreService) GetStoreBranches(storeId int64) ([]*models.Branch, error) {
	var branches []*models.Branch
	_, err := s.orm.QueryTable("sucursales").
		Filter("id_negocio", storeId).
		All(&branches)
	return branches, err
}

func (s *StoreService) UpdateStoreStatus(storeId int64, status string) error {
	store := &models.Store{Id: storeId}
	if err := s.orm.Read(store); err != nil {
		return errors.New("tienda no encontrada")
	}

	store.IsActive = (status == "activo")
	_, err := s.orm.Update(store)
	return err
}
