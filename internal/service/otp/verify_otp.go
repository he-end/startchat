package serviceotp

import (
	"errors"
	"time"

	repootp "github.com/hend41234/startchat/internal/repository/repo_otp"
)

func VerifyOtp(email, otpCode string, purpose string) (bool, error) {
	// getting otp from db
	stockOtp, err := repootp.GetOtp(email, purpose)
	// fmt.Println(err.Error())
	if err != nil {
		return false, err
	}

	// if empty otp code
	if len(stockOtp.OtpCode) < 1 {
		return false, errors.New("empty")
	}

	// check verified status
	if stockOtp.Verified {
		return false, errors.New("verified")
	}

	// matching otp code stock and input users
	if stockOtp.OtpCode != otpCode {
		return false, errors.New("not match")
	}

	// check expired
	if stockOtp.ExpiresAt.Before(time.Now()) {
		return false, errors.New("expired")
	}

	// after match, and no error, delete the otp from DB or change status verify be true
	//
	// [*] delete otp code
	// ok := repootp.DeleteSpecificOtpCode(email, otpCode)
	// if !ok {
	// 	return false, errors.New("error delete otp code")
	// }
	//  [*] change status
	if ok := repootp.UpdateStatusVerify(email, stockOtp.OtpCode); !ok {
		return false, errors.New("error update status")
	}

	return true, nil
}
