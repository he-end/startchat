package register

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	authpassword "github.com/hend41234/startchat/internal/auth/passwords"
	authvalidator "github.com/hend41234/startchat/internal/auth/validator"
	"github.com/hend41234/startchat/internal/dto"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/internalutils"
	"github.com/hend41234/startchat/internal/logger"
	mdwlogger "github.com/hend41234/startchat/internal/middleware/logger"
	repootp "github.com/hend41234/startchat/internal/repository/repo_otp"
	serviceotp "github.com/hend41234/startchat/internal/service/otp"
	servicependinguser "github.com/hend41234/startchat/internal/service/pending_user"

	"go.uber.org/zap"
)

func ResgisterHandler(w http.ResponseWriter, r *http.Request) {
	// this is a second filtering of method, cause we have filtering method in /internal/router
	if r.Method != "POST" {
		resErr := httphandler.TemplateResErr(http.StatusMethodNotAllowed, "method not allowed")
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
		resErr := httphandler.TemplateResErr(http.StatusOK, "email, password, confirm_password is required")
		w.Write(resErr)
		return
	}

	{
		// validation email, password, and Matching password
		// if ok := internalutils.EmailDetetor(reqModel.Email); !ok || // validation email
		// 	!authpassword.IsValidPassword(reqModel.Password) || // validation password
		// 	!authpassword.IsValidPassword(reqModel.ConfirmPassword) || // validation confirm_password
		// 	reqModel.Password != reqModel.ConfirmPassword { // matching password

		// 	resErr := httphandler.TemplateResErr(http.StatusOK, "please check your data")
		// 	w.Write(resErr)
		// 	return
		// }
		err := authvalidator.Validate.Struct(&reqModel)
		if err != nil {
			var invalidValidationErr *validator.InvalidValidationError
			if errors.As(err, &invalidValidationErr) {
				ctx.Error("invalid validation error", zap.Error(err))
				resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong, please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
			authvalidator.ValidationError(err, r.Context())
			resErr := httphandler.TemplateResErr(http.StatusOK, "please check your input in email and password")
			w.Write(resErr)
			return
		}
	}

	var hashPassword string
	{
		hashPassword, err = authpassword.HashingPassword(reqModel.Password)
		if err != nil {
			ctx.Error("error hasing password", zap.String("email", reqModel.Email))
			resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
			w.Write(resErr)
			return
		}
	}

	{
		// check on otp_requests
		ok, err := serviceotp.OnOrderOtpAndDelete(reqModel.Email)
		if !ok {
			if err != nil {
				resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
		}
	}

	newOTP, err := repootp.NewOTP(reqModel.Email)
	if err != nil {
		ctx.Error("error generate otp code", zap.String("email", reqModel.Email))
		resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
		w.Write(resErr)
		return
	}

	{
		// check on pending user
		ok, err := servicependinguser.OnPendingAndDelete(reqModel.Email, reqModel.Password, internalutils.GetClientIP(r))
		if !ok {
			if err != nil {
				ctx.Error("error renew pending user", zap.String("email", reqModel.Email), zap.Error(err))
				resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong, please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
		}
	}
	// add pending user
	token, err := servicependinguser.NewPendingUser(reqModel.Email, hashPassword, internalutils.GetClientIP(r))
	if err != nil {
		fmt.Println(err.Error())
	}
	resData := dto.ResRegisterModel{TokenRegister: token}
	res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: 200, Data: resData}
	byteRes, _ := json.Marshal(res)

	if err := serviceotp.SendOTPWithGmail(newOTP.OtpCode, reqModel.Email); err != nil {
		ctx.Error("error sending mail", zap.String("email", reqModel.Email), zap.Error(err))
		// resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
		// w.Write(resErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteRes)
}
