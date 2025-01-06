package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) error {
	var err error
	transaction := repo.db.Begin()

	updateBlog := &models.Blog{
		Title:          requestBody.Title,
		Description:    requestBody.Description,
		Tag:            requestBody.Tag,
		BlogCategoryID: requestBody.BlogCategoryID,
	}

	// fmt.Printf("requestBody.Image: %v\n", requestBody.Image)
	// if requestBody.Image != nil {
	// 	for i := 0; i < len(requestBody.Image); i++ {
	// 		image, err := utils.UploadFile("blog", requestBody.Image[i])
	// 		if err != nil {
	// 			return err
	// 		}

	// 		blog, err := repo.GetBlogByID(requestBody.ID)
	// 		if err != nil {
	// 			transaction.Rollback()
	// 			return err
	// 		}
	// 		err = transaction.Create(&models.Image{
	// 			URL:    image.Url,
	// 			BlogID: requestBody.ID,
	// 		}).Error

	// 		if err != nil {
	// 			transaction.Rollback()
	// 			return err
	// 		}
	// 		if blog != nil {
	// 			var imageFile models.Image

	// 			err = transaction.Debug().Model(&models.Image{}).Where("url = ?", imageFile.URL).Updates(&imageFile).Error
	// 			if err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}

	// 		}

	// 	}

	// }

	err = transaction.Debug().Where("id =?", requestBody.ID).Updates(&updateBlog).Error

	if err != nil {
		transaction.Rollback()
		return err

	}

	transaction.Commit()
	return err
}
