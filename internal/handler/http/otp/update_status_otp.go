package otp

import (
	"encoding/json"
	"net/http"

	httphandler "github.com/hend41234/startchat/internal/handler/http"
	"github.com/hend41234/startchat/internal/logger"
	mdwlogger "github.com/hend41234/startchat/internal/middleware/logger"
	servicependinguser "github.com/hend41234/startchat/internal/service/pending_user"
	"go.uber.org/zap"
)

func updateStatusOtp(w http.ResponseWriter, r *http.Request, token string) {
	ctx := logger.FromContext(r.Context())
	_, err := servicependinguser.UpdateStatusPendingUsers(token)
	if err != nil {
		ctx.Error("error update status pending_users", zap.String("token_register", token))
		resErr := httphandler.TemplateRes(http.StatusInternalServerError, nil, "something went wrong, please try again later")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resErr)
		return
	}
	ctx.Info("update status pending users", zap.String("token_register", token))
	res := httphandler.BaseResponseSuccess{ID: mdwlogger.GetRequestID(r.Context()), Code: http.StatusOK, Success: true, Metadata: ""}
	byteRes, _ := json.Marshal(res)
	w.Write(byteRes)
}
