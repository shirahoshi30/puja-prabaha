package repository

import (
	"errors"
	"fmt"

	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	auth "pujaprabha/pkg/v1/auth/handler/http"

	"gorm.io/gorm"
)

func (repo *repository) ForgotPassword(req presenter.ForgotPasswordReq) error {
	var err error
	errMap := make(map[string]string)

	user, err := GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errMap["email"] = fmt.Errorf("user with email: '%s' already exists", req.Email).Error()
			return err
		}
	}
	hashPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	err = repo.db.Model(models.User{}).Where("id = ?", user.ID).Update("password", hashPassword).Error
	if err != nil {
		return err
	}

	return nil
}
