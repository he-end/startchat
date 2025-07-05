package otp

import "time"

type ReqVerifyOTPModel struct {
	Otp           string `json:"otp"`
	TokenRegister string `json:"token_register"`
}

type ResVerifyOtpModel struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Success   bool      `json:"success"`
}
