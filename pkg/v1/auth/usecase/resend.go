package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ResendUserOTP(data presenter.UserOTPRequest) error {

	resend := uCase.repo.ResendUserOTP(data)

	return resend

}
