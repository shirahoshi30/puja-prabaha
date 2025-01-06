package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteOrderProduct(id uint) (*models.OrderProduct, error) {
	var err error

	_, err = repo.GetOrderProductsByID(id)
	if err != nil {
		return nil, err
	}

	if err = repo.db.Where("id = ?", id).Delete(&models.OrderProduct{}).Error; err != nil {
		return nil, err
	}
	return nil, err
}
