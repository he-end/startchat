package serviceotp

import (
	"errors"
	"log"
	"sc/internal/repository/repootp"
	"time"
)

func VerifyOtp(email, otpCode string) (bool, error) {
	stockOtp, err := repootp.GetOtp(email)
	if err != nil {

	}
	if len(stockOtp.OtpCode) < 1 {
		log.Println("")
	}

	if stockOtp.ExpiresAt.Before(time.Now()) {
		return false, errors.New("expired")
	}

	if stockOtp.OtpCode != otpCode {
		return false, errors.New("not match")
	}
	ok := repootp.DeleteSpecificOtpCode(email, otpCode)
	if !ok {
		return false, errors.New("error delete otp code")
	}

	return true, nil
}
