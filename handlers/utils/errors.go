package utils

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func BadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Bad Request",
		Message: message,
	}
}
func NotFound(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Not Found",
		Message: message,
	}
}

func ServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Internal Server Error",
		Message: message,
	}
}
