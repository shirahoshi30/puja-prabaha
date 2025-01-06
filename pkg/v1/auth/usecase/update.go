package usecase

import (
	"fmt"
	"pujaprabha/internal/presenter"
	"pujaprabha/pkg/v1/auth/repository"
)

func (uCase *usecase) UpdateUser(data presenter.UserCreateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)
	userID, err := repository.GetUserByID(data.ID)

	if err != nil {
		errMap["email"] = err.Error()
		return errMap
	}

	userEmail, err := repository.GetUserByEmail(data.Email)
	if err != nil {
		if userID.ID != userEmail.ID {
			errMap["email"] = fmt.Errorf("user with email: '%s' already exists", data.Email).Error()
			return errMap
		}
	}

	if err = uCase.repo.UpdateUser(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
