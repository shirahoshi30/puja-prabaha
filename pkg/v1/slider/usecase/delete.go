package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteSlider(id uint) (*models.Slider, error) {

	products, err := uCase.repo.DeleteSlider(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
