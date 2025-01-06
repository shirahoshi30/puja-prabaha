package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteCategory(id uint) (*models.Category, error) {

	category, err := uCase.repo.DeleteCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
