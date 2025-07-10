package serviceotp

import (
	"errors"
	"time"

	repootp "github.com/hend41234/startchat/internal/repository/repo_otp"
)

func VerifyOtp(email, otpCode string) (bool, error) {
	// getting otp from db
	stockOtp, err := repootp.GetOtp(email)
	if err != nil {
		return false, nil
	}
	// if nothing otp code
	if len(stockOtp.OtpCode) < 1 {
		return false, nil
	}

	// check expired
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
