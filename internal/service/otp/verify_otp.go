package serviceotp

import (
	"errors"
	"log"
	"sc/internal/repository/repootp"
	"time"
)

func VerifyOtp(email, otpCode, token string) (bool, error) {
	// getting otp from db
	stockOtp, err := repootp.GetOtp(email)
	if err != nil {

	}
	// if nothing otp code
	if len(stockOtp.OtpCode) < 1 {
		log.Println("")
	}

	// check verify
	if stockOtp.ExpiresAt.Before(time.Now()) {
		return false, errors.New("expired")
	}

	// matching otp code stock and input users
	if stockOtp.OtpCode != otpCode {
		return false, errors.New("not match")
	}

	// after match, and no error, delete the otp from DB
	ok := repootp.DeleteSpecificOtpCode(email, otpCode)
	if !ok {
		return false, errors.New("error delete otp code")
	}
	return true, nil
}
