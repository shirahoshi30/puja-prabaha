package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type SliderUsecase interface {
	CreateSlider(data presenter.SliderCreateUpdatePresenter) error
	UpdateSlider(requestBody presenter.SliderCreateUpdatePresenter) (*models.Slider, map[string]string)
	DeleteSlider(id uint) (*models.Slider, error)
	ListSlider() ([]presenter.SliderListPresenter, error)
}

type SliderRepository interface {
	CreateSlider(data presenter.SliderCreateUpdatePresenter) (*models.Slider, error)
	UpdateSlider(requestBody presenter.SliderCreateUpdatePresenter) error
	GetSliderByID(i uint) (*models.Slider, error)
	DeleteSlider(id uint) (*models.Slider, error)
	ListSlider() ([]models.Slider, error)
}
