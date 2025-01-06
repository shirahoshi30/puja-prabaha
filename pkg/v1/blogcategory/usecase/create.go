package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateBlogCategory(data presenter.BlogCategoryPresenter) (*models.BlogCategory, error) {

	blogCategory, err := uCase.repo.CreateBlogCategory(data)
	if err != nil {
		return nil, err
	}
	return blogCategory, nil

}
