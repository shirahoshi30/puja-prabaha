package repository

import (
	"errors"
	"fmt"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"time"

	"gorm.io/gorm"
)

func (repo *repository) VerifyOTP(req presenter.VerifyOTPPresenter) (bool, error) {
	var err error
	var userOTP models.UserOTP

	user, err := GetUserByEmail(req.Email)
	if err != nil {
		return false, err
	}

	err = repo.db.Model(&models.UserOTP{}).Where("user_id = ?", user.ID).First(&userOTP).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("OTP for email: '%s' not found", req.Email)
		}
		return false, err
	}

	if req.OTP != userOTP.OTP {

		return false, fmt.Errorf("wrong OTP")
	}

	if userOTP.ExpireTime.Unix() < time.Now().Unix() {

		return false, fmt.Errorf("expired OTP")
	}

	return true, nil
}
