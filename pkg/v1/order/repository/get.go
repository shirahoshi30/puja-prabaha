package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order

	err := repo.db.Where("id = ?", id).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}
