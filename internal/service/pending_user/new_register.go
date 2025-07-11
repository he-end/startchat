package servicependinguser

import (
	authrandtoken "github.com/hend41234/startchat/internal/auth/randtoken"
	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
)

func NewPendingUser(email, password string) (tokenRegister string, err error) {
	// otp, err = repootp.NewOTP(email)
	// if err != nil {
	// 	return
	// }

	token, err := authrandtoken.GenerateSecureRandomToken()
	if err != nil {
		return
	}
	tokenRegister = authrandtoken.HashRanomToken(token, authrandtoken.KeyRandT)

	err = repopendingusers.AddPendingUser(email, password, tokenRegister)
	if err != nil {
		return
	}
	return
}

func RenewPendingUser(email, newPassword string) (tokenRegister string, err error) {
	token, err := authrandtoken.GenerateSecureRandomToken()
	if err != nil {
		return
	}
	tokenRegister = authrandtoken.HashRanomToken(token, authrandtoken.KeyRandT)
	_, err = repopendingusers.UpdateOldRecord(email, newPassword, tokenRegister)
	if err != nil {
		return
	}
	return
}
