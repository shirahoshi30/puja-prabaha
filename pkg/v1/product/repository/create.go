package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateProduct(data presenter.ProductCreation) (*models.Product, error) {
	var err error
	var products models.Product
	var existingCategory models.Category

	if err := repo.db.Where("id = ?", data.CategoryID).Find(&existingCategory).Error; err != nil {
		return nil, err
	}

	product := &models.Product{
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  data.CategoryID,
		SKU:         data.SKU,
		Tags:        data.Tags,
	}
	if err := repo.db.Create(&product).Model(&products).Error; err != nil {

		return nil, err
	}

	return nil, err

}
