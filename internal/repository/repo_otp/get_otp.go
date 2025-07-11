package repootp

import (
	"database/sql"
	"errors"

	"github.com/hend41234/startchat/internal/internalutils"
	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/model"
	"github.com/hend41234/startchat/internal/repository"

	"go.uber.org/zap"
)

func GetOtp(emailOrPhone string, purpose string) (result model.OTP, err error) {
	db := repository.DB
	tx, err := db.Begin()

	if err != nil {
		logger.Error("error start transaction", zap.Error(err))
		return
	}
	var Q string

	{
		// check input email or phonr
		if email := internalutils.EmailDetetor(emailOrPhone); email {
			Q = queGetOtpFromEmail
		} else {
			Q = queGetOtpFromPhone
		}
	}

	err = tx.QueryRow(Q, emailOrPhone, purpose).Scan(
		&result.ID,
		&result.Email,
		&result.Phone,
		&result.Purpose,
		&result.OtpCode,
		&result.ExpiresAt,
		&result.Verified,
		&result.CratedAt,
		&result.VerifiedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("not found the otp code", zap.String("email", emailOrPhone), zap.Error(err))
			tx.Rollback()
			return result, nil
		}
		return result, err
	}
	defer tx.Commit()
	return
}
