package presenter

type ForgotPasswordReq struct {
	OTP             string `json:"otp"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
}
