package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) CreateSlider(data presenter.SliderCreateUpdatePresenter) (*models.Slider, error) {
	var err error

	// 	if err := repo.db.Create(&models.Slider{
	// 		Title:       data.Title,
	// 		// Image:       data.Image,
	// 		Description: data.Description,
	// 		Tag:         data.Tag,
	// 		ProductID:   data.ProductID,
	// 	}).Error; err != nil {

	// 		return nil, err
	// 	}

	// 	return nil, err

	slider := &models.Slider{
		Title:       data.Title,
		Description: data.Description,
		Tag:         data.Tag,
		ProductID:   data.ProductID,
	}
	// if data.Image != nil {
	// 	image, err := utils.UploadFile("slider", data.Image)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	// urlObject := utils.GetObjectURL(image.Url)
	// 	slider.Image = image.Key

	// }

	err = repo.db.Create(&slider).Error
	if err != nil {
		return nil, err
	}

	return nil, err

}
