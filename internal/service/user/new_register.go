package serviceuser

import (
	authrandtoken "sc/internal/auth/randtoken"
	"sc/internal/repository/repopendingusers"
)

func NewRegister(email, password, ipAddress string) (tokenRegister string, err error) {
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
