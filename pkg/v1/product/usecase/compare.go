package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) CompareProducts(id1, id2 uint) (*presenter.ProductComparePresenter, error) {
	productId1, err := uCase.DetailsOfProduct(id1)
	if err != nil {
		return nil, err
	}

	productId2, err := uCase.DetailsOfProduct(id2)

	if err != nil {
		return nil, err
	}

	response := &presenter.ProductComparePresenter{
		Product1: *productId1,
		Product2: *productId2,
	}
	return response, err
}
