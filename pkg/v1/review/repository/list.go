package repository

import (
	"math"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListReview(req presenter.ReviewList) ([]models.Review, float64, error) {
	var review []models.Review
	var err error

	if req.ProductID != 0 {
		err = repo.db.Model(&models.Review{}).Where("product_id = ?", req.ProductID).Find(&review).Error
		if err != nil {
			return nil, 0, err
		}
		totalPage := math.Ceil(float64(len(review)) / float64(req.PageSize))

		return review, totalPage, err
	}
	totalPage := utils.GetTotalPage(models.Review{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}

	err = repo.db.Debug().Model(models.Review{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&review).Order("id").Error

	return review, totalPage, err
}
