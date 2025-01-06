package usecase

import (
	"pujaprabha/internal/presenter"
	auth "pujaprabha/pkg/v1/auth/handler/http"
	"pujaprabha/pkg/v1/auth/repository"
	"pujaprabha/utils"
)

func (uCase *usecase) Login(user *presenter.LoginRequest) (presenter.LoginResponse, map[string]string) {
	errMap := make(map[string]string)
	var response presenter.LoginResponse

	result, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		errMap["error"] = "invalid username"
		return response, errMap
	}
	if !auth.CheckPasswordH(user.Password, result.Password) {
		errMap["error"] = "invalid password"
		return response, errMap
	}

	accessToken, err := utils.CreateToken(result.ID, result.Username)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	refreshToken, err := utils.RefreshToken(result.ID, result.Username)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	response = presenter.LoginResponse{
		Refresh: refreshToken,
		Access:  accessToken,
		User: presenter.UserLoginResponse{
			ID:          result.ID,
			FirstName:   result.FirstName,
			LastName:    result.LastName,
			Username:    result.Username,
			Email:       result.Email,
			PhoneNumber: result.PhoneNumber,
			Role:        string(result.Role),
		},
	}
	return response, errMap
}
