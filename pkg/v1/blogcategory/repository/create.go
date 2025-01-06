package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateBlogCategory(data presenter.BlogCategoryPresenter) (*models.BlogCategory, error) {
	var err error

	transaction := repo.db.Begin()

	blogCategory := &models.BlogCategory{
		Name: data.Name,
	}
	if err := transaction.Create(&blogCategory).Model(&blogCategory).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return nil, err

}
