package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type BlogCategoryUsecase interface {
	CreateBlogCategory(data presenter.BlogCategoryPresenter) (*models.BlogCategory, error)
	UpdateBlogCategory(requestBody presenter.BlogCategoryPresenter) (*models.BlogCategory, map[string]string)
	ListBlogCategory(requestBody presenter.BlogCategoryListReq) ([]models.BlogCategory, int, error)
	DeleteBlogCategory(id uint) (*models.BlogCategory, error)
}

type BlogCategoryRepository interface {
	CreateBlogCategory(data presenter.BlogCategoryPresenter) (*models.BlogCategory, error)
	GetBlogCategoryByID(i uint) (*models.BlogCategory, error)
	UpdateBlogCategory(data presenter.BlogCategoryPresenter) error
	ListBlogCategory(requestBody presenter.BlogCategoryListReq) ([]models.BlogCategory, float64, error)
	DeleteBlogCategory(id uint) (*models.BlogCategory, error)
}
