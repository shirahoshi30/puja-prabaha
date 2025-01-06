package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) ListSlider() ([]models.Slider, error) {

	var slider []models.Slider

	err := repo.db.Find(&slider).Error
	return slider, err
}
