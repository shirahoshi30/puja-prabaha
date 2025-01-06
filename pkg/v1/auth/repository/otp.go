package repository

import (
	"errors"
	"fmt"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func (repo *repository) CheckUserOTP(id uint) (*models.UserOTP, error) {
	var userOTP *models.UserOTP

	err := repo.db.Model(&models.UserOTP{}).Where("user_id = ?", id).First(&userOTP).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return userOTP, nil
}

func (repo *repository) ResendUserOTP(data presenter.UserOTPRequest) error {
	var err error

	generateOTP := rand.Intn(9000) + 1000
	genOTP := strconv.Itoa(generateOTP)
	fmt.Printf("genOTP: %v\n", genOTP)

	user, err := GetUserByEmail(data.Email)
	if err != nil {

		return err

	}

	userOTP, err := repo.CheckUserOTP(user.ID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if userOTP != nil {
		return repo.db.Model(&models.UserOTP{}).Where("user_id = ?", user.ID).Updates(&models.UserOTP{
			OTP:        genOTP,
			ExpireTime: time.Now().Add(time.Minute * 2),
		}).Error
	}
	err = repo.db.Create(&models.UserOTP{
		UserID:     user.ID,
		OTP:        genOTP,
		ExpireTime: time.Now().Add(time.Minute * 2),
	}).Error

	return err
}
