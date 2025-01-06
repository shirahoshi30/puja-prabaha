package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo      interfaces.OrderProductRepository
	orderRepo interfaces.OrderRepository
}

func New(repo interfaces.OrderProductRepository, orderRepo interfaces.OrderRepository) *usecase {
	return &usecase{
		repo:      repo,
		orderRepo: orderRepo,
	}
}
