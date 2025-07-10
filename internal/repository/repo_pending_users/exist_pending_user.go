package repopendingusers

import (
	"github.com/hend41234/startchat/internal/repository"
)

func PendingUserExist(email string) (bool, error) {
	tx, err := repository.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return false, err
	}
	{
		var result bool
		err = tx.QueryRow(queExistPendigUser, email).Scan(&result)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		// log.Println("user exist")
		return result, nil
	}
}
