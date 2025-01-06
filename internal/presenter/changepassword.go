package presenter

type ChangePasswordPresenter struct {
	UserID          uint   `json:"userId"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
