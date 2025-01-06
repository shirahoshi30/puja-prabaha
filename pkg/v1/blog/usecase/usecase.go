package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo             interfaces.BlogRepository
	blogCategoryRepo interfaces.BlogCategoryRepository
}

func New(repo interfaces.BlogRepository, blogCategoryRepo interfaces.BlogCategoryRepository) *usecase {
	return &usecase{
		repo:             repo,
		blogCategoryRepo: blogCategoryRepo,
	}
}
