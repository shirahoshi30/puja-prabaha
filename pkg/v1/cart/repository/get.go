package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetCartByID(id uint) (*models.Cart, error) {

	var cart models.Cart
	err := repo.db.Where("id = ?", id).First(&cart).Error
	if err != nil {

		return nil, err
	}
	return &cart, err
}
