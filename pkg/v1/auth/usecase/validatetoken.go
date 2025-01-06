package usecase

import (
	"encoding/json"
	"pujaprabha/internal/presenter"
	"pujaprabha/pkg/v1/auth/repository"
	"pujaprabha/utils"

	"github.com/gofiber/fiber/v2"
)

func (uCase *usecase) ValidateToken(c *fiber.Ctx) (*presenter.UserResponse, error) {
	token := utils.ExtractToken(c)
	claims, err := utils.ParseToken(token, "access")

	if err != nil {
		return nil, err
	}

	var userp presenter.UserResponse
	user, err := repository.GetUserByID(claims.ID)

	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(user)
	json.Unmarshal(data, &userp)
	return &userp, err
}
