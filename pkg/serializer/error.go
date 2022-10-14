package serializer

import "net/http"

func NewHttpError(code int, exception string, requestID string) *Error {
	return NewError(code, http.StatusText(code), exception, requestID)
}

func NotFound(requestID string) *Error {
	return NewHttpError(http.StatusNotFound, "Resource Not Found", requestID)
}

func BadRequest(requestID string) *Error {
	return NewHttpError(http.StatusBadRequest, "Bad Request", requestID)
}

func Unauthorized(requestID string) *Error {
	return NewHttpError(http.StatusUnauthorized, "Unauthorized", requestID)
}

func Forbidden(requestID string) *Error {
	return NewHttpError(http.StatusForbidden, "Access Denied", requestID)
}
