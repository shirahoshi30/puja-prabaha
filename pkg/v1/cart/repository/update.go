package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) UpdateCart(cartID, quantity uint) error {

	transaction := repo.db.Begin()

	var err error

	if err = transaction.Model(&models.Cart{}).Where("id = ?", cartID).Update("quantity", quantity).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err

}
