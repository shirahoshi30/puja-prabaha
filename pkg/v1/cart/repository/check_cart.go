package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) CheckCart(userID, variantID uint) (*models.Cart, error) {

	var cart models.Cart

	err := repo.db.Where("varient_id = ? AND user_id = ?", variantID, userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, err

}
