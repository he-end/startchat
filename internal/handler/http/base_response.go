package http

type BaseResponseError struct {
	Code   int         `json:"code"`
	Detail interface{} `json:"message"`
}

type BaseResponseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
