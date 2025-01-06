package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) UpdateReview(requestBody models.Review) (*models.Review, map[string]string) {
	var err error
	errMap := make(map[string]string)

	reviewID, err := uCase.repo.GetReviewByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateReview(&requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return reviewID, errMap
}
