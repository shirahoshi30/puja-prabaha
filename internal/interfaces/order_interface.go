package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type OrderUsecase interface {
	CreateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, error)
	ListOrder(req presenter.OrderListReq) ([]presenter.OrderListPresenter, int, error)
	UpdateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, map[string]string)
	DeleteOrder(id uint) (*models.Order, error)
}

type OrderRepository interface {
	CreateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, error)
	ListOrder(req presenter.OrderListReq) ([]models.Order, float64, error)
	UpdateOrder(data presenter.OrderCreateUpdatePresenter) error
	GetOrderByID(id uint) (*models.Order, error)
	DeleteOrder(id uint) (*models.Order, error)
}
