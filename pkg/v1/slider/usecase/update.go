package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) UpdateSlider(requestBody presenter.SliderCreateUpdatePresenter) (*models.Slider, map[string]string) {
	var err error
	errMap := make(map[string]string)

	sliderID, err := uCase.repo.GetSliderByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateSlider(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return sliderID, errMap
}
