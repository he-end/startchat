package authvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/hend41234/startchat/internal/dto"
	"github.com/hend41234/startchat/internal/logger"
	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
	serviceotp "github.com/hend41234/startchat/internal/service/otp"
	"go.uber.org/zap"
)

func OtpPurposeValidation(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(dto.Purpose)
	switch value {
	case dto.Register, dto.Login, dto.ForgotPassword, dto.DeleteAccount:
		return true
	default:
		return false
	}
}

func ValidatorOTPRequest(sl validator.StructLevel) {
	otpReq := sl.Current().Interface().(dto.ReqVerifyOTPModel)
	regTokenStock, err := repopendingusers.GetPendiguser3(otpReq.TokenRegister)
	if err != nil {
		logger.Error("error getting data pending user", zap.String("token_register", otpReq.TokenRegister), zap.Error(err))
		sl.ReportError(otpReq.TokenRegister, "TokenRegister", "token_register", "notfound", "")
	}

	// if otpReq.Purpose.String() == "unknown" {
	// 	logger.Error("purpose is unknown", zap.String("token_register", otpReq.TokenRegister))
	// 	sl.ReportError(otpReq.Purpose, "Purpose", "purpose", "purpose", "validation purpose error")
	// 	return
	// }
	if ok, err := serviceotp.VerifyOtp(regTokenStock.Email, otpReq.Otp, string(otpReq.Purpose)); !ok {
		if err != nil {
			if err.Error() == "expired" {
				logger.Info("verify otp failed, expired", zap.String("email", regTokenStock.Email), zap.String("otp", otpReq.Otp), zap.String("token_register", otpReq.TokenRegister), zap.String("purpose", string(otpReq.Purpose)))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "expired", "")
				return
			} else if err.Error() == "not match" {
				logger.Info("verify otp failed, not match", zap.String("email", regTokenStock.Email), zap.String("otp", otpReq.Otp), zap.String("token_register", otpReq.TokenRegister), zap.String("purpose", string(otpReq.Purpose)))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "notmatch", "")
				return
			} else if err.Error() == "error update status" {
				logger.Error("error update status", zap.String("token_register", otpReq.TokenRegister), zap.String("otp", otpReq.Otp), zap.Error(err))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "update_status_error", "")
				return
			} else if err.Error() == "empty" {
				logger.Info("code otp not found", zap.String("email", regTokenStock.Email), zap.String("token", regTokenStock.Token))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "notfound", "")
				return
			} else if err.Error() == "verified" {
				logger.Info("otp code is verified, otp is no longer valid", zap.String("otp_code", otpReq.Otp), zap.String("email", regTokenStock.Email))
				sl.ReportError(otpReq.Otp, "Otp", "otp", "verified", "")
				return
			}
			logger.Error("error verify otp", zap.String("email", regTokenStock.Email), zap.String("otp", otpReq.Otp), zap.Error(err))
		}
	}
}
