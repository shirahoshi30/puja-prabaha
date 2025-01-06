package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdatedVarientProduct(requestBody presenter.UpdateVarient) (*models.VarientProduct, map[string]string) {
	var err error
	errMap := make(map[string]string)

	varientID, err := uCase.repo.GetVarientProductByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdatedVarientProduct(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return varientID, errMap
}
