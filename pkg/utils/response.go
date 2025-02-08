package utils

import "net/http"

type ResponseStruct struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

type PaginatedResponseStruct struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
	TotalCount int         `json:"total_count"`
}

func SuccessResponse(message string, data interface{}) ResponseStruct {
	return ResponseStruct{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
}

func ErrorResponse(statusCode int, message string, err interface{}) ResponseStruct {
	return ResponseStruct{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}
}

func PaginatedResponse(message string, data interface{}, page, limit, totalPages, totalCount int) PaginatedResponseStruct {
	return PaginatedResponseStruct{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
}
