package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteReview(id uint) (*models.Review, error) {
	review, err := uCase.repo.DeleteReview(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}
