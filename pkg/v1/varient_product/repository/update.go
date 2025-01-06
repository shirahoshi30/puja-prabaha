package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdatedVarientProduct(requestBody presenter.UpdateVarient) error {

	transaction := repo.db.Begin()
	updateVarientProduct := &models.VarientProduct{
		Size:            requestBody.Size,
		Color:           requestBody.Color,
		ProductID:       uint(requestBody.ProductID),
		Price:           requestBody.Price,
		DiscountedPrice: requestBody.DiscountedPrice,
	}
	var err error

	if err = transaction.Where("id = ?", requestBody.ID).Updates(&updateVarientProduct).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err

}
