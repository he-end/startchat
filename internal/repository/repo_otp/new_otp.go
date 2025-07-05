package repootp

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hend41234/startchat/internal/internalutils"
	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/model"
	"github.com/hend41234/startchat/internal/repository"

	"go.uber.org/zap"
)

func generateNewOtp() string {
	max := 999999
	min := 100000
	newNumber := rand.Intn(max-min+1) + min
	otp := strconv.Itoa(newNumber)
	return otp
}

func NewOTP(emailOrPhone string) (result model.OTP, err error) {
	db := repository.DB

	newOTP := generateNewOtp()

	result.OtpCode = newOTP                            // result otp code
	result.ExpiresAt = time.Now().Add(time.Minute * 5) // result expires

	if email := internalutils.EmailDetetor(emailOrPhone); email {
		result.Email = repository.NullString(emailOrPhone)
		tx, err := db.Begin()
		if err != nil {
			logger.Error("error transaction DB", zap.Error(err))
			return result, err
		}
		_, err = tx.Exec(queInsertOtpFromEmail, result.Email, result.OtpCode, result.ExpiresAt)
		if err != nil {
			tx.Rollback()
			logger.Error("error insert into otp_request", zap.String("email", string(result.Email)), zap.Error(err))
			return result, err
		}
		if err := tx.Commit(); err != nil {
			logger.Error("failed to commit transaction", zap.String("email", string(result.Email)))
			return result, err
		}
	} else {
		result.Phone = repository.NullString(emailOrPhone)
		// TODO: handle phone-based OTP in future

	}
	return
}
