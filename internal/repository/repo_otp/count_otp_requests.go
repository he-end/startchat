package repootp

import "github.com/hend41234/startchat/internal/repository"

func CountOtp(email string) (int, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return 0, err
	}
	var total int
	err = tx.QueryRow(queCountOtpRequest, email).Scan(&total)
	if err != nil {
		defer tx.Rollback()
		return 0, err
	}
	defer tx.Commit()
	return total, nil
}
