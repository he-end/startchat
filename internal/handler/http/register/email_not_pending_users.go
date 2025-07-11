package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hend41234/startchat/internal/dto"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/logger"
	mdwlogger "github.com/hend41234/startchat/internal/middleware/logger"
	"github.com/hend41234/startchat/internal/model"
	serviceotp "github.com/hend41234/startchat/internal/service/otp"
	servicependinguser "github.com/hend41234/startchat/internal/service/pending_user"
	"go.uber.org/zap"
)

func NotPendingUsers(w http.ResponseWriter, r *http.Request, reqModel dto.ReqRegisterModel, hashPassword string, newOtp model.OTP) {
	ctx := logger.FromContext(r.Context())
	// add pending user
	token, err := servicependinguser.NewPendingUser(reqModel.Email, hashPassword)
	if err != nil {
		fmt.Println(err.Error())
		ctx.Error("error create new pending_users", zap.String("email", reqModel.Email), zap.Error(err))
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong, please try again later")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resErr)
		return
	}
	if err := serviceotp.SendOTPWithGmail(newOtp.OtpCode, reqModel.Email); err != nil {
		ctx.Error("error sending mail", zap.String("email", reqModel.Email), zap.Error(err))
		// resErr := httphandler.TemplateResErr(http.StatusInternalServerError, "something went wrong")
		// w.Write(resErr)
		return
	}
	resData := dto.ResRegisterModel{TokenRegister: token}
	res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: 200, Success: true, Metadata: resData}
	byteRes, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(byteRes)
}
