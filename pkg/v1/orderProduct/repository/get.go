package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetOrderProductsByID(id uint) (*models.OrderProduct, error) {
	var orderProducts models.OrderProduct

	err := repo.db.Where("id = ?", id).First(&orderProducts).Error
	if err != nil {
		return nil, err
	}
	return &orderProducts, err
}
