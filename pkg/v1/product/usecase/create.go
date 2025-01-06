package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateProduct(data presenter.ProductCreation) (*models.Product, error) {

	_, err := uCase.categoryRepo.GetCategoryByID(uint(data.CategoryID))
	if err != nil {
		return nil, err
	}

	products, err := uCase.repo.CreateProduct(data)
	if err != nil {
		return nil, err
	}
	return products, nil

}
