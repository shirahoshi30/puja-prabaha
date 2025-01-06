package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, map[string]string) {
	var err error
	errMap := make(map[string]string)

	orderID, err := uCase.repo.GetOrderByID(data.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateOrder(data)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return orderID, errMap

}
