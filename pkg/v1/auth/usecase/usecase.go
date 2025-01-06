package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.UserRepository
}

func New(repo interfaces.UserRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
