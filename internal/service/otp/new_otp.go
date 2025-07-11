package serviceotp

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/hend41234/startchat/internal/model"
	repootp "github.com/hend41234/startchat/internal/repository/repo_otp"
)

func generateNewOtp() string {
	max := 999999
	min := 100000
	newNumber := rand.Intn(max-min+1) + min
	otp := strconv.Itoa(newNumber)
	return otp
}

func NewOtp(email string, purpose string) (resultOtp model.OTP, err error) {
	if t, err := repootp.CountOtp(email); err != nil {
		return resultOtp, err
	} else if t > 5 {
		return resultOtp, errors.New("too many request")
	}
	newOTP := generateNewOtp()

	resultOtp, err = repootp.NewOTP(email, newOTP, purpose)
	return
}
