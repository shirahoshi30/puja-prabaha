package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type OrderProductUsecase interface {
	CreateOrderProduct(data presenter.OrderProductCreateListPresenter) (*models.OrderProduct, error)
	ListOrderProducts(req presenter.OrderProductListReq) ([]models.OrderProduct, int, error)
	DeleteOrderProduct(id uint) (*models.OrderProduct, error)
}

type OrderProductRepository interface {
	CreateOrderProduct(data presenter.OrderProductCreateListPresenter) (*models.OrderProduct, error)
	ListOrderProducts(requestBody presenter.OrderProductListReq) ([]models.OrderProduct, float64, error)
	GetOrderProductsByID(id uint) (*models.OrderProduct, error)
	DeleteOrderProduct(id uint) (*models.OrderProduct, error)
}
