package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) AddToCart(data presenter.AddToCart) (*models.Cart, error) {
	var err error
	var carts models.Cart

	transaction := repo.db.Begin()

	cart := &models.Cart{
		VarientID: data.VarientID,
		UserID:    data.UserID,
		Quantity:  data.Quantity,
	}
	if err := transaction.Create(&cart).Model(&carts).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return nil, err

}
