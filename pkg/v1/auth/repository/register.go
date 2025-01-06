package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) RegisterUser(user models.User) error {

	transaction := repo.db.Begin()

	err := transaction.Where("id=?", user.ID).Create(&user).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil

}
