package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) CreateReview(data *models.Review) error {
	var err error

	review := uCase.repo.CreateReview(data)
	if err != nil {
		return nil
	}
	return review
}
