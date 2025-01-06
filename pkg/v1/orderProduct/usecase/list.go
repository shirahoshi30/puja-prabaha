package usecase

import (
	"encoding/json"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListOrderProducts(req presenter.OrderProductListReq) ([]models.OrderProduct, int, error) {
	data, totalPage, err := uCase.repo.ListOrderProducts(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allOrderProduct []models.OrderProduct
	op, err := json.Marshal(data)
	if err != nil {
		return nil, int(totalPage), err
	}
	if err = json.Unmarshal(op, &allOrderProduct); err != nil {
		return nil, int(totalPage), err
	}

	return allOrderProduct, int(totalPage), nil
}
