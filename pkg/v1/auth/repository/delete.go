package repository

import "pujaprabha/internal/domain/models"

// function to delete a user
func (repo *repository) DeleteUser(id uint) (*models.User, error) {
	var err error
	transaction := repo.db.Begin()

	_, err = GetUserByID(id)

	if err != nil {
		return nil, err
	}

	if err = transaction.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return nil, err

}
