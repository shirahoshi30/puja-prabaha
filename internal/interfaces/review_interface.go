package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type ReviewUsecase interface {
	CreateReview(data *models.Review) error
	UpdateReview(requestBody models.Review) (*models.Review, map[string]string)
	DeleteReview(id uint) (*models.Review, error)
	ListReview(req presenter.ReviewList) ([]models.Review, int, error)
}

type ReviewRepository interface {
	CreateReview(data *models.Review) error
	GetReviewByID(id uint) (*models.Review, error)
	UpdateReview(data *models.Review) error
	DeleteReview(id uint) (*models.Review, error)
	ListReview(req presenter.ReviewList) ([]models.Review, float64, error)
	GetReviewByProductID(id uint) (uint, error)
}
