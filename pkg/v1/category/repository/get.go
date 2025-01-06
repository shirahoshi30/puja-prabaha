package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetCategoryByID(i uint) (*models.Category, error) {

	var category models.Category
	err := repo.db.Where("id = ?", i).First(&category).Error
	if err != nil {

		return nil, err
	}

	return &category, err
}
