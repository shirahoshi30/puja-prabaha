package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteCategory(id uint) (*models.Category, error) {
	transaction := repo.db.Begin()
	var err error

	_, err = repo.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	if err = transaction.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return nil, err
}
