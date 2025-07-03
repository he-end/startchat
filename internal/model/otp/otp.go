package modelotp

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

type ReqOTP struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type ResOTP struct {
	OTPCode   string    `json:"otp_code"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	ExpiresAt time.Time `json:"expires_at"`
}
