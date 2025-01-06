package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ListCart(uid uint) (*presenter.CartDetailPresenter, error) {
	carts, err := uCase.repo.ListCart(uint(uid))
	if err != nil {
		return nil, err
	}
	var grandTotal float32
	var allCart []presenter.CartListPresenter

	for _, cart := range carts {
		// var product models.Product

		variant, err := uCase.variantRepo.GetVarientProductByID(cart.VarientID)
		if err != nil {
			return nil, err
		}

		productName, err := uCase.variantRepo.GetProductNameWithProductID(variant.ProductID)
		if err != nil {
			return nil, err
		}

		totalPrice := cart.Quantity * uint(variant.Price)
		allCart = append(allCart, presenter.CartListPresenter{
			ID:          cart.ID,
			Quantity:    cart.Quantity,
			ProductName: productName,
			TotalPrice:  float32(totalPrice),
		})
		grandTotal += float32(totalPrice)

	}

	response := &presenter.CartDetailPresenter{
		Carts:      allCart,
		GrandTotal: grandTotal,
	}

	return response, nil
}
