package usecase

import (
	"encoding/json"
	"pujaprabha/internal/presenter"
)

func (uCase *usecase) ListUsers(req presenter.UserListReq) ([]presenter.UserListPresenter, int, error) {
	user, totalPage, err := uCase.repo.ListUsers(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allUsers []presenter.UserListPresenter
	u, err := json.Marshal(user)
	if err != nil {
		return nil, int(totalPage), err
	}
	if err = json.Unmarshal(u, &allUsers); err != nil {
		return nil, int(totalPage), err
	}

	return allUsers, int(totalPage), nil
}
