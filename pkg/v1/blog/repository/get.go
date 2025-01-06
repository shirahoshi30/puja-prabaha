package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetBlogByID(i uint) (*models.Blog, error) {

	var blog models.Blog
	err := repo.db.Where("id = ?", i).First(&blog).Error
	if err != nil {

		return nil, err
	}

	return &blog, err
}
