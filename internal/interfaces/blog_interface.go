package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type BlogUsecase interface {
	CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error)
	ListBlog(req presenter.BlogListReq) ([]presenter.BlogListPresenter, int, error)
	DetailsOfBlog(id uint) (*presenter.BlogDetailsPresenter, error)
	UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) (*models.Blog, map[string]string)
	DeleteBlog(id uint) (*models.Blog, error)
}

type BlogRepository interface {
	CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error)
	ListBlog(req presenter.BlogListReq) ([]models.Blog, float64, error)
	DetailsOfBlog(id uint) (*models.Blog, error)
	GetBlogByID(i uint) (*models.Blog, error)
	UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) error
	DeleteBlog(id uint) (*models.Blog, error)
}
