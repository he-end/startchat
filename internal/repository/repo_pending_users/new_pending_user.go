package repopendingusers

import (
	"time"

	"github.com/hend41234/startchat/internal/repository"
)

func AddPendingUser(email, password, ipAddress, token string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(queAddPendingUser, email, password, ipAddress, token, time.Now().Add(time.Minute*30))
	if err != nil {
		tx.Rollback()
		return err
	}
	defer tx.Commit()
	return nil
}
