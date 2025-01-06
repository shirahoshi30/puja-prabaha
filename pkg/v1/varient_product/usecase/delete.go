package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteVarientProduct(id uint) (*models.VarientProduct, error) {

	products, err := uCase.repo.DeleteVarientProduct(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
