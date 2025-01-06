package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) ListCart(uid uint) ([]models.Cart, error) {

	var carts []models.Cart
	err := repo.db.Where("user_id = ?", uid).Find(&carts).Error
	return carts, err
}
