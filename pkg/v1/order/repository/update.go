package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdateOrder(data presenter.OrderCreateUpdatePresenter) error {
	var err error

	updateOrder := &models.Order{
		Status: data.Status,
	}

	if err = repo.db.Where("id = ?", data.ID).Updates(&updateOrder).Error; err != nil {
		return err
	}
	return err
}
