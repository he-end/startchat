package register

import (
	"net/http"

	"github.com/hend41234/startchat/internal/dto"
	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/logger"
	"github.com/hend41234/startchat/internal/model"
	serviceotp "github.com/hend41234/startchat/internal/service/otp"
	"go.uber.org/zap"
)

func generateNewOtp(reqModel dto.ReqRegisterModel, w http.ResponseWriter, r *http.Request) model.OTP {
	ctx := logger.FromContext(r.Context())
	newOTP, err := serviceotp.NewOtp(reqModel.Email, "register")
	if err != nil {
		if err.Error() == "too many request" {
			// too many request
			ctx.Warn("to many request OTP code", zap.String("email", reqModel.Email))
			resErr := httphandler.TemplateRes(http.StatusTooManyRequests, nil, "too many request, please try again later")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write(resErr)
			return model.OTP{}
		}
		// any error
		ctx.Error("error generate otp code", zap.String("email", reqModel.Email), zap.Error(err))
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong")
		w.Write(resErr)
		return model.OTP{}
	}
	return newOTP
}
