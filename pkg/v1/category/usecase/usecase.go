package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.CategoryRepository
}

func New(repo interfaces.CategoryRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
