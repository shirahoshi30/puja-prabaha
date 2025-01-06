package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdatedProduct(requestBody presenter.UpdateProduct) error {

	transaction := repo.db.Begin()

	updateProduct := &models.Product{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		CategoryID:  requestBody.CategoryID,
	}
	var err error

	if err = transaction.Where("id = ?", requestBody.ID).Updates(&updateProduct).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err

}
