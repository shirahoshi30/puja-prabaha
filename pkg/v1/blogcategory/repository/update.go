package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdateBlogCategory(data presenter.BlogCategoryPresenter) error {
	var err error

	updateBlogCategory := &models.BlogCategory{
		Name: data.Name,
	}

	transaction := repo.db.Begin()

	if err = transaction.Where("id = ?", data.ID).Updates(&updateBlogCategory).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil

}
