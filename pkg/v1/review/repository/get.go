package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetReviewByID(id uint) (*models.Review, error) {

	var review models.Review

	err := repo.db.Where("id = ?", id).First(&review).Error
	if err != nil {

		return nil, err
	}
	return &review, err
}

func (repo *repository) GetReviewByProductID(id uint) (uint, error) {
	var product models.Product

	err := repo.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return id, err
	}
	return product.ID, err
}
