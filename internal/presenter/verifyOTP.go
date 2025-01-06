package presenter

type VerifyOTPPresenter struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
