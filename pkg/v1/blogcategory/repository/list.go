package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListBlogCategory(requestBody presenter.BlogCategoryListReq) ([]models.BlogCategory, float64, error) {

	var name []models.BlogCategory

	totalPage := utils.GetTotalPage(models.BlogCategory{}, repo.db, requestBody.PageSize)

	if totalPage < float64(requestBody.Page) {
		requestBody.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.BlogCategory{}).Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&name).Order("id").Error

	return name, totalPage, err
}
