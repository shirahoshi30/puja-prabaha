package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.OrderRepository
}

func New(repo interfaces.OrderRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
