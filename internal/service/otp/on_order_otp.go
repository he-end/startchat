package serviceotp

import (
	"github.com/hend41234/startchat/internal/logger"
	repootp "github.com/hend41234/startchat/internal/repository/repo_otp"
	"go.uber.org/zap"
)

func OnOrderOtpAndDelete(email string) (ok bool, err error) {
	ok, err = repootp.ExistOtpOrder(email)
	if !ok {
		if err != nil {
			logger.Error("error get otp", zap.String("email", email), zap.Error(err))
			return false, err
		}
		return false, nil
	}
	// deleting old otp_requests
	ok = repootp.DeleteOtpCodeWithEmail(email)
	return
}
