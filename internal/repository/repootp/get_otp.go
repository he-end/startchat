package repootp

import (
	"database/sql"
	"errors"
	"sc/internal/internalutils"
	"sc/internal/logger"
	modelotp "sc/internal/model/otp"
	"sc/internal/repository"

	"go.uber.org/zap"
)

func GetOtp(emailOrPhone string) (result modelotp.OTP, err error) {
	db := repository.DB
	tx, err := db.Begin()
	if err != nil {
		logger.Error("error start transaction", zap.Error(err))
		return
	}
	var Q string

	if email := internalutils.EmailDetetor(emailOrPhone); email {
		Q = queGetOtpFromEmail
	} else {
		Q = queGetOtpFromPhone
	}

	err = tx.QueryRow(Q, emailOrPhone).Scan(
		&result.ID,
		&result.Email,
		&result.Phone,
		&result.OtpCode,
		&result.ExpiresAt,
		&result.Verified,
		&result.CratedAt,
		&result.UpdatedAt,
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
