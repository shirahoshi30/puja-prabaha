package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteProduct(id uint) (*models.Product, error) {
	transaction := repo.db.Begin()

	_, err := repo.DetailsOfProduct(id)

	if err != nil {
		return nil, err
	}
	if err = transaction.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return nil, err

}
