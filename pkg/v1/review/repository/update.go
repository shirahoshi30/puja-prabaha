package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) UpdateReview(data *models.Review) error {

	updateReview := &models.Review{
		Rating:  data.Rating,
		Message: data.Message,
	}
	var err error

	if err = repo.db.Where("id = ?", data.ID).Updates(&updateReview).Error; err != nil {
		return err
	}

	return err
}
