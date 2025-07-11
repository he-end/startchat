package register

import (
	"encoding/json"
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

func ReRegistrationPendingUsers(w http.ResponseWriter, r *http.Request, reqModel dto.ReqRegisterModel, newOtp model.OTP, hashPassword string) {
	ctx := logger.FromContext(r.Context())

	// resending otp code
	if err := serviceotp.SendOTPWithGmail(newOtp.OtpCode, reqModel.Email); err != nil {
		ctx.Error("error sending mail", zap.String("email", reqModel.Email), zap.Error(err))
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong")
		w.Write(resErr)
		return
	}
	// new update
	token, err := servicependinguser.RenewPendingUser(reqModel.Email, hashPassword)
	if err != nil {
		ctx.Error("error re-registration user on step pending_user", zap.String("email", reqModel.Email), zap.Error(err))
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong, please try again later")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resErr)
		return
	}
	tknRes := dto.ResRegisterModel{TokenRegister: token}
	res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: http.StatusOK, Success: true, Metadata: tknRes}
	byteRes, err := json.Marshal(res)
	if err != nil {
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wong, please try agin later")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resErr)
		return
	}
	w.Write(byteRes)
	return
}
