package repository

import "pujaprabha/internal/domain/models"

// function to delete a product
func (repo *repository) DeleteVarientProduct(id uint) (*models.VarientProduct, error) {
	transaction := repo.db.Begin()
	var err error

	_, err = repo.GetVarientProductByID(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Debug().Where("id = ?", id).Delete(&models.VarientProduct{}).Error

	if err != nil {

		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return nil, err

}
