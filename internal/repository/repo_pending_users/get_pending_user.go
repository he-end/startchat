package repopendingusers

import (
	"database/sql"
	"errors"

	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/model"
	"github.com/hend41234/startchat/internal/repository"

	"go.uber.org/zap"
)

func GetPendingUser(email, token string) (pu model.PendingUserModel, err error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return pu, err
	}
	err = tx.QueryRow(queGetPendingUser, email, token).Scan(&pu.Email, &pu.Password, &pu.IpAddress, &pu.CreatedAt, &pu.ExpiresAt)
	if err != nil {
		defer tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("not found the otp code", zap.String("email", email), zap.Error(err))
			return pu, nil
		}
		return pu, err
	}
	tx.Commit()
	return
}
