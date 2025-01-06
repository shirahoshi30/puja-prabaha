package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.VarientProductRepository
}

func New(repo interfaces.VarientProductRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
