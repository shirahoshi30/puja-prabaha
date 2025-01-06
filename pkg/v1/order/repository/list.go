package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListOrder(req presenter.OrderListReq) ([]models.Order, float64, error) {
	var orders []models.Order
	totalPage := utils.GetTotalPage(models.Order{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.Order{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&orders).Order("id").Error

	return orders, totalPage, err
}
