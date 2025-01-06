package models

import "time"

type UserOTP struct {
	BaseModel
	UserID     uint      `json:"userId"`
	OTP        string    `json:"otp"`
	ExpireTime time.Time `json:"expireTime"`
}
