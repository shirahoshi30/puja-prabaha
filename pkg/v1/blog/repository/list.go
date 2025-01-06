package repository

import (
	"math"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListBlog(req presenter.BlogListReq) ([]models.Blog, float64, error) {

	var blog []models.Blog

	if req.BlogCategoryID != 0 {

		err := repo.db.Debug().Model(models.Blog{}).Where("category_id = ?", req.BlogCategoryID).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&blog).Order("id").Error

		totalPage := math.Ceil(float64(len(blog)) / float64(req.PageSize))

		return blog, totalPage, err

	}

	totalPage := utils.GetTotalPage(models.Blog{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.Blog{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&blog).Order("id").Error

	return blog, totalPage, err
}
