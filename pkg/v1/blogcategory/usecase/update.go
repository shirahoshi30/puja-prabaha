package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdateBlogCategory(requestBody presenter.BlogCategoryPresenter) (*models.BlogCategory, map[string]string) {
	var err error
	errMap := make(map[string]string)

	blogCategoryID, err := uCase.repo.GetBlogCategoryByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}
	err = uCase.repo.UpdateBlogCategory(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return blogCategoryID, errMap
}
