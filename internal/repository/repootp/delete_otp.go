package repootp

import (
	"sc/internal/logger"
	"sc/internal/repository"

	"go.uber.org/zap"
)

func DeleteSpecificOtpCode(email, otpCode string) bool {
	db := repository.DB
	tx, err := db.Begin()
	if err != nil {
		logger.Error("error start transaction", zap.Error(err))
		return false
	}

	_, err = tx.Exec(queDeleteSpecificOtpCode, email, otpCode)
	if err != nil {
		logger.Error("failed delete otp code", zap.String("email", email), zap.Error(err))
		tx.Rollback()
		return false
	}

	defer tx.Commit()

	return true

}

func DeleteOtpCodeWithEmail(email string) bool {
	tx, err := repository.DB.Begin()
	if err != nil {
		logger.Error("error start transaction", zap.Error(err))
		return false
	}

	_, err = tx.Exec(queDeleteOtpFromEmail, email)
	if err != nil {
		logger.Error("error delete otp code with email", zap.String("email", email), zap.Error(err))
		tx.Rollback()
		return false
	}
	defer tx.Commit()
	return true
}
