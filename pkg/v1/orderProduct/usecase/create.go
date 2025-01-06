package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateOrderProduct(data presenter.OrderProductCreateListPresenter) (*models.OrderProduct, error) {
	_, err := uCase.orderRepo.GetOrderByID(data.OrderID)
	if err != nil {
		return nil, err
	}

	orderProduct, err := uCase.repo.CreateOrderProduct(data)
	if err != nil {
		return nil, err
	}
	return orderProduct, err
}
