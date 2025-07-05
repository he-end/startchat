package register

import (
	"encoding/json"
	"fmt"
	"net/http"
	authpassword "sc/internal/auth/passwords"
	httphandler "sc/internal/handler/http"
	"sc/internal/internalutils"
	"sc/internal/logger"
	mdwlogger "sc/internal/middleware/logger"
	"sc/internal/repository/repootp"
	serviceotp "sc/internal/service/otp"
	serviceuser "sc/internal/service/user"

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

	var reqModel ReqRegisterModel // model request
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
		if ok := internalutils.EmailDetetor(reqModel.Email); !ok || // validation email
			!authpassword.IsValidPassword(reqModel.Password) || // validation password
			!authpassword.IsValidPassword(reqModel.ConfirmPassword) || // validation confirm_password
			reqModel.Password != reqModel.ConfirmPassword { // matching password

			resErr := httphandler.TemplateResErr(http.StatusOK, "please check your data")
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

	fmt.Println("password : ", hashPassword)
	newOTP, err := repootp.NewOTP(reqModel.Email)
	if err != nil {
		ctx.Error("error generate otp code", zap.String("email", reqModel.Email))
		resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
		w.Write(resErr)
		return
	}

	if err := serviceotp.SendOTPWithGmail(newOTP.OtpCode, reqModel.Email); err != nil {
		ctx.Error("error sending mail", zap.String("email", reqModel.Email), zap.Error(err))
		resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
		w.Write(resErr)
		return
	}

	// add pending user
	token, err := serviceuser.NewRegister(reqModel.Email, reqModel.Password, internalutils.GetClientIP(r))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token, newOTP.CratedAt)
	resData := ResRegisterModel{TokenRegister: token}
	res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: 200, Data: resData}
	byteRes, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(byteRes)
}
