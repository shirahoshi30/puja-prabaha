package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) DetailsOfBlog(id uint) (*presenter.BlogDetailsPresenter, error) {

	blog, err := uCase.repo.DetailsOfBlog(id)
	if err != nil {
		return nil, err
	}

	detail := &presenter.BlogDetailsPresenter{
		ID:          blog.ID,
		Tag:         blog.Tag,
		Title:       blog.Title,
		Description: blog.Description,
		ReadCount:   uint(blog.ReadCount),
		Image:       blog.Image,
	}

	blogCategory, _ := uCase.blogCategoryRepo.GetBlogCategoryByID(uint(blog.BlogCategoryID))

	blogCategoryMap := make(map[string]interface{})

	blogCategoryMap["id"] = blogCategory.ID
	blogCategoryMap["name"] = blogCategory.Name
	detail.BlogCategory = blogCategoryMap

	return detail, nil
}
