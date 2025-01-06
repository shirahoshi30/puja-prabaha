package interfaces

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

type UserUsecase interface {
	Login(*presenter.LoginRequest) (presenter.LoginResponse, map[string]string)
	DeleteUser(id uint) (*models.User, error)
	RegisterUser(user models.User) map[string]string
	UpdateUser(data presenter.UserCreateRequest) map[string]string
	ListUsers(req presenter.UserListReq) ([]presenter.UserListPresenter, int, error)
	ForgotPassword(data presenter.ForgotPasswordReq) (*models.UserOTP, error)
	VerifyOTP(otpRequest presenter.VerifyOTPPresenter) (bool, error)
	ResendUserOTP(data presenter.UserOTPRequest) error
	ChangePassword(data presenter.ChangePasswordPresenter) error
}


type UserRepository interface {
	DeleteUser(id uint) (*models.User, error)
	RegisterUser(user models.User) error
	UpdateUser(data presenter.UserCreateRequest) error
	ListUsers(req presenter.UserListReq) ([]models.User, float64, error)
	CheckUserOTP(id uint) (*models.UserOTP, error)
	ForgotPassword(req presenter.ForgotPasswordReq) error
	ResendUserOTP(data presenter.UserOTPRequest) error
	VerifyOTP(req presenter.VerifyOTPPresenter) (bool, error)
	ChangePassword(req presenter.ChangePasswordPresenter) error
}
