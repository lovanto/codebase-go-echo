package utils

import "net/http"

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

// Success response
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
}

// Error response
func ErrorResponse(statusCode int, message string, err interface{}) Response {
	return Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}
}
