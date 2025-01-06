package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteProduct(id uint) (*models.Product, error) {

	products, err := uCase.repo.DeleteProduct(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
