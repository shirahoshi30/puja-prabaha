package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ListBlog(req presenter.BlogListReq) ([]presenter.BlogListPresenter, int, error) {
	blog, totalPage, err := uCase.repo.ListBlog(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allBlog []presenter.BlogListPresenter
	for _, blogs := range blog {

		// url := utils.GetObjectURL(blogs.Image)
		// if err != nil {
		// 	return nil, int(totalPage), err
		// }

		allBlog = append(allBlog, presenter.BlogListPresenter{
			ID:    blogs.ID,
			Title: blogs.Title,
			Tag:   blogs.Tag,
		})

	}

	return allBlog, int(totalPage), err
}
