package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdatedProduct(requestBody presenter.UpdateProduct) (*models.Product, map[string]string) {
	var err error
	errMap := make(map[string]string)

	productID, err := uCase.repo.DetailsOfProduct(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdatedProduct(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return productID, errMap
}
