package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteUser(id uint) (*models.User, error) {

	products, err := uCase.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
