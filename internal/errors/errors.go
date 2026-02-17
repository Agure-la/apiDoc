package errors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) HTTPStatus() int {
	switch e.Code {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

const (
	ErrNotFound    = "NOT_FOUND"
	ErrConflict    = "CONFLICT"
	ErrBadRequest  = "BAD_REQUEST"
	ErrInternal    = "INTERNAL_ERROR"
	ErrValidation  = "VALIDATION_ERROR"
)

func NewAPIError(code, message, details string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func NotFound(resource string) *APIError {
	return NewAPIError(ErrNotFound, fmt.Sprintf("%s not found", resource), "")
}

func Conflict(resource string, details string) *APIError {
	return NewAPIError(ErrConflict, fmt.Sprintf("%s conflict", resource), details)
}

func BadRequest(message string) *APIError {
	return NewAPIError(ErrBadRequest, message, "")
}

func InternalError(message string) *APIError {
	return NewAPIError(ErrInternal, "Internal server error", message)
}

func ValidationError(message string) *APIError {
	return NewAPIError(ErrValidation, message, "")
}
