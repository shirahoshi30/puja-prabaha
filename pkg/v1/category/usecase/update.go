package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdatedCategory(requestBody presenter.CategoryPresenter) (*models.Category, map[string]string) {
	var err error
	errMap := make(map[string]string)

	categoryID, err := uCase.repo.GetCategoryByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}
	err = uCase.repo.UpdatedCategory(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return categoryID, errMap
}
