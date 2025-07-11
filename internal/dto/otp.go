package dto

import (
	"time"
)

type Purpose string

const (
	Register       = "register"
	Login          = "login"
	ForgotPassword = "forgot_password"
	DeleteAccount  = "delete_account"
)

type ReqVerifyOTPModel struct {
	Otp           string  `json:"otp"`
	TokenRegister string  `json:"token_register"`
	Purpose       Purpose `json:"purpose" validate:"required,otp_purpose"`
}

type ResVerifyOtpModel struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Success   bool      `json:"success"`
}
