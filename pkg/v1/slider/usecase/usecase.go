package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.SliderRepository
}

func New(repo interfaces.SliderRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
