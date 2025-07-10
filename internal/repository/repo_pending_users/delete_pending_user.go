package repopendingusers

import "github.com/hend41234/startchat/internal/repository"

func DeletePendingUsers(email, token string) error {
	tx, err := repository.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}
	_, err = tx.Exec(queDeletePendingUser, email, token)
	if err != nil {
		defer tx.Rollback()
		return err
	}
	return nil
}

func DeletePendingUsers2(email string) error {
	tx, err := repository.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}
	_, err = tx.Exec(queDeletePendingUser2, email)
	if err != nil {
		defer tx.Rollback()
		return err
	}
	return nil
}
