package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo         interfaces.ProductRepository
	categoryRepo interfaces.CategoryRepository
}

func New(repo interfaces.ProductRepository, categoryRepo interfaces.CategoryRepository) *usecase {
	return &usecase{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}
