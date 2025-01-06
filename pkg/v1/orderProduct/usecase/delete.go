package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteOrderProduct(id uint) (*models.OrderProduct, error) {
	orderProduct, err := uCase.repo.DeleteOrderProduct(id)
	if err != nil {
		return nil, err
	}
	return orderProduct, err
}
