package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteOrder(id uint) (*models.Order, error) {
	var err error

	_, err = repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	if err = repo.db.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		return nil, err
	}
	return nil, err
}
