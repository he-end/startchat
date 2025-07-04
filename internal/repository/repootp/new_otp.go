package repootp

import (
	"math/rand"
	"sc/internal/internalutils"
	"sc/internal/logger"
	modelotp "sc/internal/model/otp"
	"sc/internal/repository"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func generateNewOtp() string {
	max := 999999
	min := 100000
	newNumber := rand.Intn(max-min+1) + min
	otp := strconv.Itoa(newNumber)
	return otp
}

func NewOTP(emailOrPhone string) (resOTP modelotp.ResOTP, err error) {
	db := repository.DB

	newOTP := generateNewOtp()

	resOTP.OTPCode = newOTP
	resOTP.ExpiresAt = time.Now().Add(time.Minute * 5)

	if email := internalutils.EmailDetetor(emailOrPhone); email {
		resOTP.Email = emailOrPhone
		tx, err := db.Begin()
		if err != nil {
			logger.Error("error transaction DB", zap.Error(err))
			return resOTP, err
		}
		_, err = tx.Exec(queInsertOtpFromEmail, resOTP.Email, resOTP.OTPCode, resOTP.ExpiresAt)
		if err != nil {
			tx.Rollback()
			logger.Error("error insert into otp_request", zap.String("email", resOTP.Email), zap.Error(err))
			return resOTP, err
		}
		if err := tx.Commit(); err != nil {
			logger.Error("failed to commit transaction", zap.String("email", resOTP.Email))
			return resOTP, err
		}
	} else {
		resOTP.Phone = emailOrPhone
		// TODO: handle phone-based OTP in future

	}
	return
}
