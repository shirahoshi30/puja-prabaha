package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) RemoveCart(id uint) (*models.Cart, error) {

	cart, err := uCase.repo.RemoveCart(id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
