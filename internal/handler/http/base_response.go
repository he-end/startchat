package httphandler

type BaseResponseError struct {
	Code        int         `json:"code"`
	ErrorDetail interface{} `json:"error_detail,omitempty"`
	Message     interface{} `json:"message,omitempty"`
}

type BaseResponseSuccess struct {
	ID       string      `json:"id"`
	Code     int         `json:"code"`
	Success  bool        `json:"success"`
	Metadata interface{} `json:"metadata,omitempty"`
}
