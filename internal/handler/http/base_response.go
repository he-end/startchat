package httphandler

type BaseResponseError struct {
	Code        int         `json:"code"`
	ErrorDetail interface{} `json:"error_detail"`
}

type BaseResponseSuccess struct {
	ID   string      `json:"id"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
