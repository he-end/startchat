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
	"github.com/hend41234/startchat/internal/logger"
	"go.uber.org/zap"
)

// verify OTP, POST Method
func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	ctx := logger.FromContext(r.Context())
	var reqModel dto.ReqVerifyOTPModel
	err := json.NewDecoder(r.Body).Decode(&reqModel)
	if err != nil {
		resErr := httphandler.TemplateRes(http.StatusOK, nil, "otp, token_register is required")
		w.WriteHeader(http.StatusOK)
		w.Write(resErr)
		ctx.Warn("body request error", zap.Error(err))
		return
	}

	ipClient := internalutils.GetClientIP(r)
	if ipClient == "" {
		resErr := httphandler.TemplateRes(http.StatusBadRequest, nil, "something went wrong")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resErr)
		return
	}

	{
		// validate input and verify otp code
		err := authvalidator.Validate.Struct(reqModel)
		if err != nil {
			var invalidValidationErr *validator.InvalidValidationError
			if errors.As(err, &invalidValidationErr) {
				resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong, please try again later")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resErr)
				return
			}
			valErr := authvalidator.ValidationError(err, r.Context())
			// e, _ := json.Marshal(valErr)
			for _, v := range valErr {
				if v["tag"] == "verified" {
					resErr := httphandler.TemplateRes(http.StatusBadRequest, nil, "bad request")
					w.WriteHeader(http.StatusBadRequest)
					w.Write(resErr)
					return
				}
			}
			resErr := httphandler.TemplateRes(http.StatusOK, nil, valErr)
			w.Write(resErr)
			return
		}
	}

	{
		// TODO : migration data user from pending_users
		// migration data user

	}

	{
		// update status pending users
		updateStatusOtp(w, r, reqModel.TokenRegister)
		return
	}

	// res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: http.StatusOK, Success: true, Metadata: ""}
	// byteRes, _ := json.Marshal(res)
	// w.Write(byteRes)
}
