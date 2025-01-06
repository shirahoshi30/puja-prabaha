package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateVarientProduct(data presenter.VarientProductCreation) (*models.VarientProduct, error) {

	products, err := uCase.repo.CreateVarientProduct(data)
	if err != nil {
		return nil, err
	}
	return products, nil

}
