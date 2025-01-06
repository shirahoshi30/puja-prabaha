package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type CartUsecase interface {
	AddToCart(data presenter.AddToCart) error
	UpdateCart(requestBody presenter.UpdateCart) (*models.Cart, map[string]string)
	ListCart(uid uint) (*presenter.CartDetailPresenter, error)
	RemoveCart(id uint) (*models.Cart, error)
}

type CartRepository interface {
	AddToCart(data presenter.AddToCart) (*models.Cart, error)
	UpdateCart(cartID, quantity uint) error
	ListCart(uid uint) ([]models.Cart, error)
	GetCartByID(id uint) (*models.Cart, error)
	RemoveCart(id uint) (*models.Cart, error)
	CheckCart(userID, variantID uint) (*models.Cart, error)
}
