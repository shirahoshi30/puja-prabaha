package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteBlogCategory(id uint) (*models.BlogCategory, error) {
	transaction := repo.db.Begin()
	var err error

	_, err = repo.GetBlogCategoryByID(id)

	if err != nil {
		return nil, err
	}

	if err = transaction.Where("id = ?", id).Delete(&models.BlogCategory{}).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return nil, err
}
