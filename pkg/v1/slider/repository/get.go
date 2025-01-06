package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetSliderByID(i uint) (*models.Slider, error) {

	var slider models.Slider
	err := repo.db.Where("id = ?", i).First(&slider).Error
	if err != nil {

		return nil, err
	}

	return &slider, err
}
