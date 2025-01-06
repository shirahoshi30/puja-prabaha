package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type VarientProductUsecase interface {
	CreateVarientProduct(data presenter.VarientProductCreation) (*models.VarientProduct, error)
	ListVarient(req presenter.VariantListReq) ([]presenter.VarientListRes, int, error)
	UpdatedVarientProduct(requestBody presenter.UpdateVarient) (*models.VarientProduct, map[string]string)
	DeleteVarientProduct(id uint) (*models.VarientProduct, error)
}

type VarientProductRepository interface {
	CreateVarientProduct(data presenter.VarientProductCreation) (*models.VarientProduct, error)
	ListVarient(req presenter.VariantListReq) ([]models.VarientProduct, float64, error)
	UpdatedVarientProduct(requestBody presenter.UpdateVarient) error
	DeleteVarientProduct(id uint) (*models.VarientProduct, error)
	GetVarientProductByID(i uint) (*models.VarientProduct, error)
	GetProductNameWithProductID(id uint) (string, error)
}
