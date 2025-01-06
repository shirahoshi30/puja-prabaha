package repository

import (
	"errors"
	"pujaprabha/internal/config"
	"pujaprabha/internal/domain/models"

	"gorm.io/gorm"
)

func (repo *repository) DetailsOfBlog(id uint) (*models.Blog, error) {
	db, err := config.ConfigDb()
	if err != nil {
		return nil, err
	}

	var blog *models.Blog
	err = db.Where("id=?", id).First(&blog).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	err = db.Model(&models.Blog{}).Where("id = ?", id).Update("read_count", blog.ReadCount+1).Error

	if err != nil {
		return nil, err
	}

	return blog, err
}
