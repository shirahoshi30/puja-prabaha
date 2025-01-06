package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteBlog(id uint) (*models.Blog, error) {

	var err error

	_, err = repo.GetBlogByID(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Debug().Where("id = ?", id).Delete(&models.Blog{}).Error

	if err != nil {

		return nil, err
	}

	return nil, err

}
