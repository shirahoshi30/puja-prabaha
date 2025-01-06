package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteSlider(id uint) (*models.Slider, error) {

	var err error

	_, err = repo.GetSliderByID(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Debug().Where("id = ?", id).Delete(&models.Slider{}).Error

	if err != nil {

		return nil, err
	}

	return nil, err

}
