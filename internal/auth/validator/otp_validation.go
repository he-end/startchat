package authvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/hend41234/startchat/internal/dto"
	"github.com/hend41234/startchat/internal/logger"
	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
	serviceotp "github.com/hend41234/startchat/internal/service/otp"
	"go.uber.org/zap"
)

func ValidatorOTPRequest(sl validator.StructLevel) {
	otpReq := sl.Current().Interface().(dto.ReqVerifyOTPModel)
	regTokenStock, err := repopendingusers.GetPendiguser3(otpReq.TokenRegister)
	if err != nil {
		logger.Error("error getting data pending user", zap.String("token_register", otpReq.TokenRegister), zap.Error(err))
		sl.ReportError(otpReq.TokenRegister, "TokenRegister", "token_register", "tokenregister", "")
	}

	if ok, err := serviceotp.VerifyOtp(regTokenStock.Email, otpReq.Otp); !ok {
		if err != nil {
			if err.Error() == "expired" || err.Error() == "not match" {
				logger.Info("verify otp failed", zap.String("email", regTokenStock.Email))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "otp", "otp expired or not match")
			}
			logger.Error("error verify otp", zap.String("email", regTokenStock.Email), zap.String("otp", otpReq.Otp), zap.Error(err))
		}
		sl.ReportError(otpReq.Otp, "Otp", "otp", "otp", "")
	}
}
