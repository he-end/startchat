package repopendingusers

import (
	"fmt"

	"github.com/hend41234/startchat/internal/model"
	"github.com/hend41234/startchat/internal/repository"
)

// func GetPendingUser(email, token string) (pu model.PendingUserModel, err error) {
// 	tx, err := repository.DB.Begin()
// 	if err != nil {
// 		return pu, err
// 	}
// 	err = tx.QueryRow(queGetPendingUser, email, token).Scan(&pu.Id, &pu.Email, &pu.Password, &pu.IpAddress, &pu.CreatedAt, &pu.ExpiresAt)
// 	if err != nil {
// 		defer tx.Rollback()
// 		if errors.Is(err, sql.ErrNoRows) {
// 			logger.Error("not found the otp code", zap.String("email", email), zap.Error(err))
// 			return pu, nil
// 		}
// 		return pu, err
// 	}
// 	tx.Commit()
// 	return
// }

// func GetPendinguser2(ipAddress string, regToken string) (pu model.PendingUserModel, err error) {
// 	tx, err := repository.DB.Begin()
// 	defer tx.Commit()
// 	if err != nil {
// 		return
// 	}

// 	err = tx.QueryRow(queGetPendingUser2, ipAddress, regToken).Scan(&pu.Id, &pu.Email, &pu.Password, &pu.IpAddress, &pu.CreatedAt, &pu.ExpiresAt)
// 	if err != nil {
// 		defer tx.Rollback()
// 		return pu, err
// 	}
// 	return
// }

func GetPendiguser3(regToken string) (pu model.PendingUserModel, err error) {
	tx, err := repository.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return
	}

	err = tx.QueryRow(queGetPendingUser3, regToken).Scan(&pu.ID, &pu.Email, &pu.Password, &pu.IpAddress, &pu.Token, &pu.CreatedAt, &pu.ExpiresAt)
	if err != nil {
		defer tx.Rollback()
		fmt.Println(err.Error())
		return pu, err
	}
	return pu, nil
}
