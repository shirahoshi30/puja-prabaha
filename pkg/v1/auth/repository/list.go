package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListUsers(req presenter.UserListReq) ([]models.User, float64, error) {

	var users []models.User
	totalPage := utils.GetTotalPage(models.User{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err := repo.db.Debug().Model(models.User{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&users).Order("id").Error

	return users, totalPage, err
}
