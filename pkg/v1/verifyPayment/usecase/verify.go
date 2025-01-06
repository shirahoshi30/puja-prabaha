package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) VerifyPayment(req presenter.VerifyPaymentPresenter) (*models.Payment, error) {
	var err error
	verify, _ := uCase.repo.VerifyPayment(req)
	if err != nil {
		return nil, err
	}

	return verify, err
}
