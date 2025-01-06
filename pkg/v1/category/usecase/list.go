package usecase

import (
	"encoding/json"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListCategory(requestBody presenter.CategoryListReq) ([]models.Category, int, error) {
	name, totalPage, err := uCase.repo.ListCategory(requestBody)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allCategory []models.Category
	pp, err := json.Marshal(name)
	if err != nil {
		return nil, int(totalPage), err
	}
	if err = json.Unmarshal(pp, &allCategory); err != nil {
		return nil, int(totalPage), err
	}

	return allCategory, int(totalPage), nil
}
