package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetBlogCategoryByID(i uint) (*models.BlogCategory, error) {

	var blogCategory models.BlogCategory
	err := repo.db.Where("id = ?", i).First(&blogCategory).Error
	if err != nil {

		return nil, err
	}

	return &blogCategory, err
}
