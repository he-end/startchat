package model

import (
	"time"

	"github.com/hend41234/startchat/internal/repository"
)

type OTP struct {
	ID         string                `json:"id"`
	Email      repository.NullString `json:"email"`
	Phone      repository.NullString `json:"phone"`
	Purpose    string                `json:"purpose"`
	OtpCode    string                `json:"otp_code"`
	ExpiresAt  time.Time             `json:"expires_at"`
	Verified   bool                  `json:"verified"`
	CratedAt   time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
	VerifiedAt *time.Time            `json:"verified_at"`
}

// type ReqOTP struct {
// 	Email string `json:"email,omitempty"`
// 	Phone string `json:"phone,omitempty"`
// }

// type ResOTP struct {
// 	OTPCode   string    `json:"otp_code"`
// 	Email     string    `json:"email,omitempty"`
// 	Phone     string    `json:"phone,omitempty"`
// 	ExpiresAt time.Time `json:"expires_at"`
// }
