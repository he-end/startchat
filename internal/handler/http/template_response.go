package httphandler

import (
	"encoding/json"
	"sc/internal/logger"

	"go.uber.org/zap"
)

func TemplateResErr(code int, msg string) (byteRes []byte) {
	res := BaseResponseError{Code: code, ErrorDetail: msg}
	byteRes, err := json.Marshal(res)
	if err != nil {
		logger.Error("decode template response error", zap.Error(err))
		return
	}
	return
}
