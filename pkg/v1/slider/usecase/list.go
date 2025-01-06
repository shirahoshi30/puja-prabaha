package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ListSlider() ([]presenter.SliderListPresenter, error) {
	slider, err := uCase.repo.ListSlider()
	if err != nil {
		return nil, err
	}
	var allSlider []presenter.SliderListPresenter
	for _, sliders := range slider {

		// url := utils.GetObjectURL(sliders.Image)
		// if err != nil {
		// 	return nil, err
		// }

		allSlider = append(allSlider, presenter.SliderListPresenter{
			ID:          sliders.ID,
			Title:       sliders.Title,
			Description: sliders.Description,
			Tag:         sliders.Tag,
			ProductID:   sliders.ProductID,
			// Image:       url,
		})

	}

	return allSlider, err
}
