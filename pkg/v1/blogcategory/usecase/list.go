package usecase

import (
	"encoding/json"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListBlogCategory(requestBody presenter.BlogCategoryListReq) ([]models.BlogCategory, int, error) {
	name, totalPage, err := uCase.repo.ListBlogCategory(requestBody)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allBlogCategory []models.BlogCategory
	pp, err := json.Marshal(name)
	if err != nil {
		return nil, int(totalPage), err
	}
	if err = json.Unmarshal(pp, &allBlogCategory); err != nil {
		return nil, int(totalPage), err
	}

	return allBlogCategory, int(totalPage), nil
}
