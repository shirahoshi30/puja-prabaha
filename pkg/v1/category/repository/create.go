package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateCategory(data presenter.CategoryPresenter) (*models.Category, error) {
	var err error

	transaction := repo.db.Begin()

	category := &models.Category{
		Name: data.Name,
	}
	if err := transaction.Create(&category).Model(&category).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return nil, err

}
