package repository

import (
	"math"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListAllProduct(req presenter.ProductListReq) ([]models.Product, float64, error) {

	var product []models.Product

	if req.CategoryID != 0 {

		err := repo.db.Debug().Model(models.Product{}).Where("category_id = ?", req.CategoryID).Preload("VarientProducts").Preload("Reviews").Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&product).Order("id").Error

		totalPage := math.Ceil(float64(len(product)) / float64(req.PageSize))

		return product, totalPage, err

	}

	totalPage := utils.GetTotalPage(models.Product{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.Product{}).Preload("VarientProducts").Preload("Reviews").Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&product).Order("id").Error

	return product, totalPage, err

}
