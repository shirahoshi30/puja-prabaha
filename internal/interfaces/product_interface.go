package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type ProductUsecase interface {
	ListAllProduct(req presenter.ProductListReq) ([]presenter.ProductListPresenter, int, error)
	DetailsOfProduct(id uint) (*presenter.ProductDetailsPresenter, error)
	UpdatedProduct(requestBody presenter.UpdateProduct) (*models.Product, map[string]string)
	CreateProduct(data presenter.ProductCreation) (*models.Product, error)
	DeleteProduct(id uint) (*models.Product, error)
	CompareProducts(id1, id2 uint) (*presenter.ProductComparePresenter, error)
}

type ProductRepository interface {
	ListAllProduct(req presenter.ProductListReq) ([]models.Product, float64, error)
	DetailsOfProduct(id uint) (*models.Product, error)
	UpdatedProduct(requestBody presenter.UpdateProduct) error
	CreateProduct(data presenter.ProductCreation) (*models.Product, error)
	DeleteProduct(id uint) (*models.Product, error)
}
