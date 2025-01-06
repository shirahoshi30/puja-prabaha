package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type CategoryUsecase interface {
	CreateCategory(data presenter.CategoryPresenter) (*models.Category, error)
	ListCategory(requestBody presenter.CategoryListReq) ([]models.Category, int, error)
	UpdatedCategory(requestBody presenter.CategoryPresenter) (*models.Category, map[string]string)
	DeleteCategory(id uint) (*models.Category, error)
}

type CategoryRepository interface {
	CreateCategory(data presenter.CategoryPresenter) (*models.Category, error)
	ListCategory(requestBody presenter.CategoryListReq) ([]models.Category, float64, error)
	GetCategoryByID(i uint) (*models.Category, error)
	UpdatedCategory(data presenter.CategoryPresenter) error
	DeleteCategory(id uint) (*models.Category, error)
}
