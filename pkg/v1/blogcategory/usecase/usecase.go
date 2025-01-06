package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.BlogCategoryRepository
}

func New(repo interfaces.BlogCategoryRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
