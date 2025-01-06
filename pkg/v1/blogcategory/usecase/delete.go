package usecase

import "pujaprabha/internal/domain/models"

func (uCase *usecase) DeleteBlogCategory(id uint) (*models.BlogCategory, error) {

	category, err := uCase.repo.DeleteBlogCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
