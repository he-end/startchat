package httphandler

import (
	"encoding/json"

	"github.com/hend41234/startchat/internal/logger"

	"go.uber.org/zap"
)

func TemplateRes(code int, msg interface{}, e interface{}) (byteRes []byte) {
	res := BaseResponseError{Code: code, ErrorDetail: e, Message: msg}
	byteRes, err := json.Marshal(res)
	if err != nil {
		logger.Error("decode template response error", zap.Error(err))
		return
	}
	return
}
