package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) AddToCart(data presenter.AddToCart) error {

	existingCart, _ := uCase.repo.CheckCart(data.UserID, data.VarientID)

	if existingCart == nil {
		_, err := uCase.repo.AddToCart(data)
		if err != nil {
			return err
		}
	} else {
		quantity := existingCart.Quantity + data.Quantity
		err := uCase.repo.UpdateCart(existingCart.ID, quantity)
		if err != nil {
			return err
		}
	}

	return nil

}
