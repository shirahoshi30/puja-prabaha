package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) DeleteReview(id uint) (*models.Review, error) {
	var err error

	_, err = repo.GetReviewByID(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Where("id = ?", id).Delete(&models.Review{}).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}
