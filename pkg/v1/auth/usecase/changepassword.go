package usecase

import (
	"errors"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ChangePassword(data presenter.ChangePasswordPresenter) error {
	var err error

	if data.OldPassword == data.NewPassword {

		return errors.New("cannot reuse old password")

	}

	if data.NewPassword != data.ConfirmPassword {

		return errors.New("passwords do not match")

	}

	err = uCase.repo.ChangePassword(data)
	if err != nil {
		return err
	}
	return err
}
