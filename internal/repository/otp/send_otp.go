package otp

import (
	"errors"
	"fmt"
	internalutils "sc/internal/utils"
)

func SendOTP(email string) error {
	if ok := internalutils.EmailDetetor(email); !ok {
		return errors.New("error location : 'otp.SendOtp'\nemail is not valid")
	}
	newOtp := generateNewOtp()
	fmt.Println(newOtp)
	return nil
}
