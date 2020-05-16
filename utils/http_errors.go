package utils

import "errors"

// ErrorResponse common struct to http error
type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// BadRequest http error
func BadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Bad Request",
		Message: message,
	}
}

// NotFound http error
func NotFound(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Not Found",
		Message: message,
	}
}

// ErrSomethingWrong for unexpected errors
var ErrSomethingWrong = errors.New("Something went wrong")

// ServerError http error
func ServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Internal Server Error",
		Message: message,
	}
}
