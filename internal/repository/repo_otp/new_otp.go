package repootp

import (
	"fmt"
	"time"

	"github.com/hend41234/startchat/internal/internalutils"
	"github.com/hend41234/startchat/internal/model"
	"github.com/hend41234/startchat/internal/repository"
)

func NewOTP(emailOrPhone string, newOTP string, purpose string) (result model.OTP, err error) {
	db := repository.DB

	result.OtpCode = newOTP                            // result otp code
	result.ExpiresAt = time.Now().Add(time.Minute * 5) // result expires

	if email := internalutils.EmailDetetor(emailOrPhone); email {
		result.Email = repository.NullString(emailOrPhone)
		tx, err := db.Begin()
		if err != nil {
			fmt.Println("================================")
			// logger.Error("error transaction DB", zap.Error(err))
			return result, err
		}
		_, err = tx.Exec(queInsertOtpFromEmail, result.Email, result.OtpCode, purpose, result.ExpiresAt)
		if err != nil {
			defer tx.Rollback()
			// logger.Error("error insert into otp_request", zap.String("email", string(result.Email)), zap.Error(err))
			return result, err
		}
		if err := tx.Commit(); err != nil {
			// logger.Error("failed to commit transaction", zap.String("email", string(result.Email)))
			return result, err
		}
	} else {
		result.Phone = repository.NullString(emailOrPhone)
		// TODO: handle phone-based OTP in future

	}
	return
}
