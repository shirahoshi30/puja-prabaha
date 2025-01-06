package usecase

import (
	"encoding/json"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListOrder(req presenter.OrderListReq) ([]presenter.OrderListPresenter, int, error) {
	order, totalPage, err := uCase.repo.ListOrder(req)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allOrders []presenter.OrderListPresenter
	o, err := json.Marshal(order)
	if err != nil {
		return nil, int(totalPage), err
	}
	if err = json.Unmarshal(o, &allOrders); err != nil {
		return nil, int(totalPage), err
	}
	return allOrders, int(totalPage), err
}
