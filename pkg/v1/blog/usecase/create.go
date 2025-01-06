package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error) {

	_, err := uCase.blogCategoryRepo.GetBlogCategoryByID(uint(data.BlogCategoryID))
	if err != nil {
		return nil, err
	}

	blog, err := uCase.repo.CreateBlog(data)
	if err != nil {
		return nil, err
	}
	return blog, err

}
