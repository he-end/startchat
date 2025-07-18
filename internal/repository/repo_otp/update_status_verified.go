package repootp

import (
	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/repository"

	"go.uber.org/zap"
)

func UpdateStatusVerify(email, otpCode string) bool {
	tx, err := repository.DB.Begin()
	if err != nil {
		// log.Println("error start transaction : ", err.Error())
		logger.Error("error start transaction DB", zap.Error(err))
		return false
	}
	res, err := tx.Exec(queUpdateStatusVerifyOtp, email, otpCode)
	if err != nil {
		logger.Error("error update status verified otp", zap.String("email", email), zap.Error(err))
		tx.Rollback()
		return false
	}

	defer tx.Commit()
	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Info("update success but no user found", zap.String("email", email))
	}
	return true
}
