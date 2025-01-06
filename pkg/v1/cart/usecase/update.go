package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdateCart(requestBody presenter.UpdateCart) (*models.Cart, map[string]string) {
	var err error
	errMap := make(map[string]string)
	cartID, err := uCase.repo.GetCartByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	// err = uCase.repo.UpdateCart(requestBody.ID, requestBody.Quantity)
	// if err != nil {
	// 	errMap["update"] = err.Error()
	// 	return nil, errMap
	// }

	if requestBody.Quantity == 0 {
		_, err = uCase.repo.RemoveCart(requestBody.ID)
		if err != nil {
			errMap["remove"] = err.Error()
			return nil, errMap
		}

	} else {
		err = uCase.repo.UpdateCart(requestBody.ID, requestBody.Quantity)
		if err != nil {
			errMap["update"] = err.Error()
			return nil, errMap
		}
	}

	return cartID, errMap
}
