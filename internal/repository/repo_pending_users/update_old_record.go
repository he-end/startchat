package repopendingusers

import (
	"time"

	"github.com/hend41234/startchat/internal/repository"
)

func UpdateOldRecord(email string, newPassword string, token string) (bool, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(queUpdateOldRecord, newPassword, time.Now().Add(time.Minute*5), token, email)
	if err != nil {
		defer tx.Rollback()
		return false, err
	}
	defer tx.Commit()
	return true, nil
}
