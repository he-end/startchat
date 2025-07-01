package model

import "time"

type OTP struct {
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	OtpCode   string    `json:"otp_code"`
	ExpiresAt time.Time `json:"expires_at"`
	Verified  bool      `json:"verified"`
	CratedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

