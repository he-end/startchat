package repopendingusers

import (
	"time"

	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/repository"
	"go.uber.org/zap"
)

func DeleterPendingUsersExpired() (bool, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Commit()
	res, err := tx.Exec(queCleanerPendingUserExpired)
	if err != nil {
		defer tx.Rollback()
		return false, err
	}
	rows, _ := res.RowsAffected()
	if rows > 0 {
		timezone, _ := time.Now().Zone()
		logger.Info("deleted pending users", zap.Int("lenght", int(rows)), zap.Time(timezone, time.Now()))
	}
	return true, nil
}

func DeleterPendingUsersVerified() (bool, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return false, err
	}
	res, err := tx.Exec(queCleanerPendingUserVerified)
	if err != nil {
		return false, err
	}
	rows, _ := res.RowsAffected()
	if rows > 0 {
		timezone, _ := time.Now().Zone()
		logger.Info("deleted pending users", zap.Int("lenght", int(rows)), zap.Time(timezone, time.Now()))
	}
	defer tx.Commit()
	return true, nil
}
