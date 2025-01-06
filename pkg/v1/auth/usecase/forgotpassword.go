package usecase

import (
	"errors"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/pkg/v1/auth/repository"
)

func (uCase *usecase) ForgotPassword(data presenter.ForgotPasswordReq) (*models.UserOTP, error) {
	var err error

	user, err := repository.GetUserByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	_, err = uCase.repo.VerifyOTP(presenter.VerifyOTPPresenter{
		Email: user.Email,
		OTP:   data.OTP,
	})
	if err != nil {
		return nil, err
	}

	if data.NewPassword != data.ConfirmPassword {

		return nil, errors.New("passwords do not match")
	}

	err = uCase.repo.ForgotPassword(data)
	if err != nil {
		return nil, err
	}
	return nil, err
}
