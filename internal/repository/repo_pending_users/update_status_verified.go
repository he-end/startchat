package repopendingusers

import (
	"github.com/hend41234/startchat/internal/repository"
)

func UpdateStatusPending(token string) (bool, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(queUpdateStatusPending, token)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	defer tx.Commit()
	return true, nil
}
