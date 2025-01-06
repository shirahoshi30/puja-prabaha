package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListCategory(requestBody presenter.CategoryListReq) ([]models.Category, float64, error) {

	var name []models.Category

	totalPage := utils.GetTotalPage(models.Category{}, repo.db, requestBody.PageSize)

	if totalPage < float64(requestBody.Page) {
		requestBody.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.Category{}).Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&name).Order("id").Error

	return name, totalPage, err
}
