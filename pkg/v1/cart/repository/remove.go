package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) RemoveCart(id uint) (*models.Cart, error) {
	transaction := repo.db.Begin()

	_, err := repo.GetCartByID(id)

	if err != nil {
		return nil, err
	}
	if err = transaction.Where("id = ?", id).Delete(&models.Cart{}).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return nil, err

}
