package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	auth "pujaprabha/pkg/v1/auth/handler/http"
)

func (repo *repository) ChangePassword(req presenter.ChangePasswordPresenter) error {
	var err error

	user, err := GetUserByID(req.UserID)
	if err != nil {
		return err
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
