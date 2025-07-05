package otp

import (
	"encoding/json"
	"fmt"
	"net/http"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/internalutils"
)

// verify OTP, POST Method
func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := logger.FromContext(r.Context())
	var reqModel ReqVerifyOTPModel
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

	fmt.Println(ipClient, reqModel.Otp, reqModel.TokenRegister)

	res := httphandler.BaseResponseSuccess{Code: http.StatusOK, Data: "success"}
	byteRes, _ := json.Marshal(res)
	w.Write(byteRes)
}
