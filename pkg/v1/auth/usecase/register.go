package usecase

import (
	"fmt"
	"pujaprabha/internal/domain/models"
	auth "pujaprabha/pkg/v1/auth/handler/http"
	"pujaprabha/pkg/v1/auth/repository"
)

func (uCase *usecase) RegisterUser(user models.User) map[string]string {
	var err error
	errMap := make(map[string]string)
	if _, err := repository.GetUserByEmail(user.Email); err == nil {
		errMap["email"] = fmt.Errorf("user with the email: '%s' already exists", user.Email).Error()
	}

	if _, err := repository.GetUserByUsername(user.Username); err == nil {
		errMap["username"] = fmt.Errorf("user with the username: '%s' already exists", user.Username).Error()
	}
	if _, err := repository.GetUserByPhone(int(user.PhoneNumber)); err == nil {
		errMap["phonenumber"] = fmt.Errorf("user with the phone number: '%d' already exists", user.PhoneNumber).Error()
	}

	if len(errMap) != 0 {
		return errMap
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	if err = uCase.repo.RegisterUser(user); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	return nil

}
