package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateOrderProduct(data presenter.OrderProductCreateListPresenter) (*models.OrderProduct, error) {
	var err error
	var order models.OrderProduct
	var existingOrder models.Order

	if err := repo.db.Where("id = ?", data.OrderID).Find(&existingOrder).Error; err != nil {
		return nil, err
	}

	orderProduct := &models.OrderProduct{
		UserID:    data.UserID,
		ProductID: data.ProductId,
		OrderID:   data.OrderID,
	}
	if err := repo.db.Create(&orderProduct).Model(&order).Error; err != nil {

		return nil, err
	}

	return nil, err

}
