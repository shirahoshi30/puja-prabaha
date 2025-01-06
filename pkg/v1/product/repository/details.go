package repository

import (
	"errors"
	"pujaprabha/internal/config"
	"pujaprabha/internal/domain/models"

	"gorm.io/gorm"
)

func (repo *repository) DetailsOfProduct(id uint) (*models.Product, error) {
	db, err := config.ConfigDb()
	if err != nil {
		return nil, err
	}

	var product *models.Product
	err = db.Debug().Where("id=?", id).Preload("VarientProducts").Preload("Reviews").First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return product, err
}
