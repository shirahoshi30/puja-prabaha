package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdatedCategory(data presenter.CategoryPresenter) error {
	var err error

	UpdatedCategory := &models.Category{
		Name: data.Name,
	}

	transaction := repo.db.Begin()

	if err = transaction.Where("id = ?", data.ID).Updates(&UpdatedCategory).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil

}
