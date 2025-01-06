package usecase

import "pujaprabha/internal/interfaces"

type usecase struct {
	repo interfaces.VerifyPaymentRepository
}

func New(repo interfaces.VerifyPaymentRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
