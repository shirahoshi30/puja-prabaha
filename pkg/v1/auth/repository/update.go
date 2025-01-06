package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

/*
updating user's information in the database
parameter
data : instance of UpdateUsersRequest struct which contains necessary information to
update a user.
*/

func (repo *repository) UpdateUser(data presenter.UserCreateRequest) error {
	var err error

	updatedUser := &models.User{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Username:    data.Username,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
	}

	transaction := repo.db.Begin()

	if err = transaction.Where("id = ?", data.ID).Updates(&updatedUser).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil

}
