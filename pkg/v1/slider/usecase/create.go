package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) CreateSlider(data presenter.SliderCreateUpdatePresenter) error {

	_, err := uCase.repo.CreateSlider(data)
	if err != nil {
		return err
	}
	return err

}
