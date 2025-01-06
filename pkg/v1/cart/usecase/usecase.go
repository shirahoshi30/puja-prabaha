package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo        interfaces.CartRepository
	variantRepo interfaces.VarientProductRepository
	productRepo interfaces.ProductRepository
}

func New(repo interfaces.CartRepository, variantRepo interfaces.VarientProductRepository, productRepo interfaces.ProductRepository) *usecase {
	return &usecase{
		repo:        repo,
		variantRepo: variantRepo,
		productRepo: productRepo,
	}
}
