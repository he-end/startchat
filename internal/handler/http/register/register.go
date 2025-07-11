package register

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	authpassword "github.com/hend41234/startchat/internal/auth/passwords"
	authvalidator "github.com/hend41234/startchat/internal/auth/validator"
	"github.com/hend41234/startchat/internal/dto"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/model"
	servicependinguser "github.com/hend41234/startchat/internal/service/pending_user"

	"go.uber.org/zap"
)

func ResgisterHandler(w http.ResponseWriter, r *http.Request) {
	// this is a second filtering of method, cause we have filtering method in /internal/router
	if r.Method != "POST" {
		resErr := httphandler.TemplateRes(http.StatusMethodNotAllowed, nil, "method not allowed")
		if len(resErr) == 0 {
			logger.Error("internal server error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(resErr)
		return
	}
	ctx := logger.FromContext(r.Context()) // ctx logger

	var reqModel dto.ReqRegisterModel // model request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&reqModel)
	if err != nil {
		resErr := httphandler.TemplateRes(http.StatusOK, nil, "email, password, confirm_password is required")
		w.Write(resErr)
		return
	}

	{
		// validation email, password, and Matching password
		// if ok := internalutils.EmailDetetor(reqModel.Email); !ok || // validation email
		// 	!authpassword.IsValidPassword(reqModel.Password) || // validation password
		// 	!authpassword.IsValidPassword(reqModel.ConfirmPassword) || // validation confirm_password
		// 	reqModel.Password != reqModel.ConfirmPassword { // matching password

		// 	resErr := httphandler.TemplateRes(http.StatusOK, "please check your data")
		// 	w.Write(resErr)
		// 	return
		// }
		err := authvalidator.Validate.Struct(&reqModel)
		if err != nil {
			var invalidValidationErr *validator.InvalidValidationError
			if errors.As(err, &invalidValidationErr) {
				ctx.Error("invalid validation error", zap.Error(err))
				resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong, please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
			authvalidator.ValidationError(err, r.Context())
			resErr := httphandler.TemplateRes(http.StatusOK, nil, "please check your input in email and password")
			w.Write(resErr)
			return
		}
	}

	var hashPassword string
	{
		hashPassword, err = authpassword.HashingPassword(reqModel.Password)
		if err != nil {
			ctx.Error("error hasing password", zap.String("email", reqModel.Email))
			resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong")
			w.Write(resErr)
			return
		}
	}

	// {
	// 	// check on otp_requests
	// 	ok, err := serviceotp.OnOrderOtpAndDelete(reqModel.Email)
	// 	if !ok {
	// 		if err != nil {
	// 			resErr := httphandler.TemplateRes(http.StatusInternalServerError, "something went wrong")
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			w.Write(resErr)
	// 			return
	// 		}
	// 	}
	// }
	var newOtp model.OTP
	var pendingUserExpired bool
	{

		// check on pending user
		ok, err := servicependinguser.CheckPendingUser(reqModel.Email)
		if !ok {
			switch err.Error() {
			case "email is already to use": // email already to use and isnt expired
				ctx.Info("new register rejected, because email is already in pending users", zap.String("email", reqModel.Email))
				resErr := httphandler.TemplateRes(http.StatusOK, "email already to use,we have sent OTP to your inbox", nil)
				w.Write(resErr)
				return
			default: // any error, may error in internal / server
				ctx.Error("check pending users error", zap.String("email", reqModel.Email), zap.Error(err))
				resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong,please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
		}

		if err != nil {
			if err.Error() == "pending expired" {
				newOtp = generateNewOtp(reqModel, w, r)
				if newOtp.OtpCode == "" {
					return
				}
				pendingUserExpired = true
			}
		}
	}

	if newOtp.OtpCode == "" {
		// generate new otp code
		newOtp = generateNewOtp(reqModel, w, r)
		if newOtp.OtpCode == "" {
			return
		}
	}

	if !pendingUserExpired {
		ctx.Info("new register, and new on pending users", zap.String("email", reqModel.Email))
		NotPendingUsers(w, r, reqModel, hashPassword, newOtp)
		return
	} else {
		ctx.Info("re-registration, and resending otp code", zap.String("email", reqModel.Email))
		ReRegistrationPendingUsers(w, r, reqModel, newOtp, hashPassword)
		return
	}
}
