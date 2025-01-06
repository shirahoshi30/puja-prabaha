package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) VerifyOTP(otpRequest presenter.VerifyOTPPresenter) (bool, error) {

	verified, err := uCase.repo.VerifyOTP(otpRequest)

	return verified, err
}
