package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateOrder(data presenter.OrderCreateUpdatePresenter) (*models.Order, error) {
	var err error
	order := &models.Order{
		UserID:      data.UserID,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		GrandTotal:  data.GrandTotal,
		PaymentMode: data.PaymentMode,
		Status:      data.Status,
	}
	if err := repo.db.Create(&order).Model(&order).Error; err != nil {
		return nil, err
	}
	return nil, err

}
