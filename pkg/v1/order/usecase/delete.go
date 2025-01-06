package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteOrder(id uint) (*models.Order, error) {
	order, err := uCase.repo.DeleteOrder(id)
	if err != nil {
		return nil, err
	}
	return order, err
}
