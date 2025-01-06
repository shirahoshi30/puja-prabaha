package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type VerifyPaymentUsecase interface {
	VerifyPayment(req presenter.VerifyPaymentPresenter) (*models.Payment, error)
}

type VerifyPaymentRepository interface {
	VerifyPayment(req presenter.VerifyPaymentPresenter) (*models.Payment, error)
}
