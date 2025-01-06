package repository

import (
	"math"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListOrderProducts(requestBody presenter.OrderProductListReq) ([]models.OrderProduct, float64, error) {

	var orderProduct []models.OrderProduct

	if requestBody.OrderID != 0 {

		err := repo.db.Debug().Model(models.OrderProduct{}).Where("order_id = ?", requestBody.OrderID).Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&orderProduct).Order("id").Error

		totalPage := math.Ceil(float64(len(orderProduct)) / float64(requestBody.PageSize))

		return orderProduct, totalPage, err

	}

	totalPage := utils.GetTotalPage(models.OrderProduct{}, repo.db, requestBody.PageSize)

	if totalPage < float64(requestBody.Page) {
		requestBody.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.OrderProduct{}).Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&orderProduct).Order("id").Error

	return orderProduct, totalPage, err

}
