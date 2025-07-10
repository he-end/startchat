package servicependinguser

import (
	authrandtoken "github.com/hend41234/startchat/internal/auth/randtoken"
	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
)

func NewPendingUser(email, password, ipAddress string) (tokenRegister string, err error) {
	// otp, err = repootp.NewOTP(email)
	// if err != nil {
	// 	return
	// }

	token, err := authrandtoken.GenerateSecureRandomToken()
	if err != nil {
		return
	}
	tokenRegister = authrandtoken.HashRanomToken(token, authrandtoken.KeyRandT)

	err = repopendingusers.AddPendingUser(email, password, ipAddress, tokenRegister)
	if err != nil {
		return
	}
	return
}

func RenewPendingUser(email, password, ipAddress string) (tokenRegister string, err error) {

	// DELETE old pending user
	if err = repopendingusers.DeletePendingUsers2("email"); err != nil {
		return
	}

	token, err := authrandtoken.GenerateSecureRandomToken()
	if err != nil {
		return
	}
	tokenRegister = authrandtoken.HashRanomToken(token, authrandtoken.KeyRandT)

	err = repopendingusers.AddPendingUser(email, password, ipAddress, tokenRegister)
	if err != nil {
		return
	}
	return
}

func OnPendingAndDelete(email, password, ipAddress string) (ok bool, err error) {
	ok, err = repopendingusers.PendingUserExist(email)
	if !ok {
		if err != nil {
			// logger.Error("error check pending user exist", zap.String("email", email))
			return false, err
		}
		return false, nil
	}

	err = repopendingusers.DeletePendingUsers2(email)
	if err != nil {
		return false, err
	}
	return
}
