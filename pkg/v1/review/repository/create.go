package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) CreateReview(data *models.Review) error {
	transaction := repo.db.Begin()

	review := &models.Review{
		Rating:    data.Rating,
		Message:   data.Message,
		ProductID: data.ProductID,
		UserID:    data.UserID,
	}
	if err := transaction.Create(&review).Model(&review).Error; err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
