package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateCategory(data presenter.CategoryPresenter) (*models.Category, error) {

	category, err := uCase.repo.CreateCategory(data)
	if err != nil {
		return nil, err
	}
	return category, nil

}
