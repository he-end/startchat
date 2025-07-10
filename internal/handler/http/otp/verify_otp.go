package otp

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	authvalidator "github.com/hend41234/startchat/internal/auth/validator"
	"github.com/hend41234/startchat/internal/dto"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/internalutils"
)

// verify OTP, POST Method
func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := logger.FromContext(r.Context())
	var reqModel dto.ReqVerifyOTPModel
	err := json.NewDecoder(r.Body).Decode(&reqModel)
	if err != nil {
		resErr := httphandler.TemplateResErr(http.StatusOK, "otp, token_register is required")
		w.WriteHeader(http.StatusOK)
		w.Write(resErr)
		// ctx.Warn("body request error", zap.Error(err))
		return
	}

	ipClient := internalutils.GetClientIP(r)
	if ipClient == "" {
		resErr := httphandler.TemplateResErr(http.StatusBadRequest, "something went wrong")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resErr)
		return
	}

	{
		// validate input
		err := authvalidator.Validate.Struct(reqModel)
		if err != nil {
			var invalidValidationErr *validator.InvalidValidationError
			if errors.As(err, &invalidValidationErr) {
				resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong, please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
			authvalidator.ValidationError(err, r.Context())
			resErr := httphandler.TemplateResErr(http.StatusOK, "make sure data is correct, please check your data")
			w.Write(resErr)
			return
		}
	}

	res := httphandler.BaseResponseSuccess{Code: http.StatusOK, Data: "success, please login"}
	byteRes, _ := json.Marshal(res)
	w.Write(byteRes)
}
