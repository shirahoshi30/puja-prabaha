package repository

import (
	"errors"
	"fmt"
	"pujaprabha/internal/config"
	"pujaprabha/internal/domain/models"

	"gorm.io/gorm"
)

// gets the data of the user corresponding the specific email
func GetUserByEmail(e string) (*models.User, error) {
	db, err := config.ConfigDb()

	var user *models.User
	if err := db.Where("email = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with the given email")
		}
		return nil, err
	}
	return user, err
}

// gets the data of the user corresponding the specific username
func GetUserByUsername(u string) (*models.User, error) {
	db, err := config.ConfigDb()

	var user *models.User
	if err := db.Where("username = ?", u).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with the given username")
		}
		return nil, err
	}
	return user, err
}

// gets the data of the user corresponding the specific id
func GetUserByID(i uint) (*models.User, error) {
	db, err := config.ConfigDb()

	var user models.User
	if err := db.Where("id = ?", i).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}

func GetUserRole(u string) (*models.User, error) {
	db, err := config.ConfigDb()

	var user models.User
	if err := db.Debug().Where(&models.User{Role: models.UserRole(u)}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return nil, err
}

func GetUserByPhone(p int) (*models.User, error) {
	db, err := config.ConfigDb()
	var user models.User

	if err := db.Where("phone_number = ?", p).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}
