package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, error) {
	order, err := uCase.repo.CreateOrder(data)
	if err != nil {
		return nil, err
	}
	return order, err
}
