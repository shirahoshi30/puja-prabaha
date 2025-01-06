package usecase

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListReview(req presenter.ReviewList) ([]models.Review, int, error) {
	review, totalPage, err := uCase.repo.ListReview(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var productReview []models.Review

	for i := 0; i < len(review); i++ {

		prodID, err := uCase.repo.GetReviewByProductID(review[i].ProductID)
		if err != nil {
			return nil, int(totalPage), err
		}
		productReview = append(productReview, models.Review{
			BaseModel: models.BaseModel{
				ID: review[i].ID,
			},
			Rating:    review[i].Rating,
			Message:   review[i].Message,
			ProductID: prodID,
			UserID:    review[i].UserID,
		})
	}

	// var reviews []models.Review
	// u, err := json.Marshal(review)
	// if err != nil {
	// 	return nil, int(totalPage), err
	// }
	// if err = json.Unmarshal(u, &reviews); err != nil {
	// 	return nil, int(totalPage), err
	// }

	// return reviews, int(totalPage), nil
	return productReview, int(totalPage), nil
}
