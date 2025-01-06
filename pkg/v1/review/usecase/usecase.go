package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.ReviewRepository
}

func New(repo interfaces.ReviewRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
