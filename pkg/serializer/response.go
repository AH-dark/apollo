package serializer

type Response struct {
	Code    int         `json:"code" xml:"Code"`
	Message string      `json:"message" xml:"Message"`
	Data    interface{} `json:"data" xml:"Data"`
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type Error struct {
	Code      int    `json:"code" xml:"Code"`
	Message   string `json:"message" xml:"Message"`
	Exception string `json:"exception" xml:"Exception"`
	RequestID string `json:"request_id" xml:"RequestID"`
}

func NewError(code int, message string, exception string, requestID string) *Error {
	return &Error{
		Code:      code,
		Message:   message,
		Exception: exception,
		RequestID: requestID,
	}
}
