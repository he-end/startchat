package register

import (
	"encoding/json"
	"fmt"
	"net/http"
	httphandler "sc/internal/handler/http"
	"sc/internal/logger"
)

func ResgisterHandler(w http.ResponseWriter, r *http.Request) {
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
	res := httphandler.BaseResponseSuccess{Code: 200, Data: "success"}
	byteRes, _ := json.Marshal(res)
	{
		fmt.Println(r.Header.Get("X-Request-ID"))

	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteRes)
}
