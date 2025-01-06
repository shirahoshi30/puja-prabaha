package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error) {
	var err error
	transaction := repo.db.Begin()
	var existingBlogCategory models.Category

	if err := repo.db.Where("id = ?", data.BlogCategoryID).Find(&existingBlogCategory).Error; err != nil {
		return nil, err
	}
	blog := &models.Blog{
		Title:          data.Title,
		Description:    data.Description,
		Tag:            data.Tag,
		BlogCategoryID: data.BlogCategoryID,
	}

	err = transaction.Create(&blog).Error
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	// if data.Image != nil {
	// 	for i := 0; i < len(data.Image); i++ {
	// 		image, err := utils.UploadFile("blog", data.Image[i])
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if err = transaction.Create(&models.Image{
	// 			URL:    image.Url,
	// 			BlogID: blog.ID,
	// 		}).Error; err != nil {
	// 			transaction.Rollback()
	// 			return nil, err
	// 		}

	// 		// blog.Image = image.Key
	// 	}

	// }
	transaction.Commit()
	return nil, err

}
