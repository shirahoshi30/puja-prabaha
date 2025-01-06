package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

// function to create a product
func (repo *repository) CreateVarientProduct(data presenter.VarientProductCreation) (*models.VarientProduct, error) {
	var err error
	var vProducts models.VarientProduct

	transaction := repo.db.Begin()

	vProduct := &models.VarientProduct{
		Size:            data.Size,
		Color:           data.Color,
		Price:           data.Price,
		DiscountedPrice: data.DiscountedPrice,
		ProductID:       uint(data.ProductID),
	}
	if err := transaction.Create(&vProduct).Model(&vProducts).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return nil, err

}
