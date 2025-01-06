package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) UpdateSlider(requestBody presenter.SliderCreateUpdatePresenter) error {
	var err error

	updateSlider := &models.Slider{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Tag:         requestBody.Tag,
	}

	// if requestBody.Image != nil {
	// 	image, err := utils.UploadFile("slider", requestBody.Image)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// urlObject := utils.GetObjectURL(image.Url)
	// 	updateSlider.Image = image.Key
	// }

	err = repo.db.Where("id = ?", requestBody.ID).Updates(&updateSlider).Error
	if err != nil {
		return err
	}
	return err

}
