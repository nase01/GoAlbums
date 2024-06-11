package helpers

import (
	"log"
	"net/http"
	"strings"
)

type ErrorDetail struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

func CustomError(err error) (ErrorResponse, int) {
	// Default error response
	statusCode := http.StatusInternalServerError
	detail := "Internal Server Error"

	log.Printf("err: ,%v", err.Error())

	errorDetail := ErrorDetail{
		Status: statusCode,
		Detail: detail,
	}

	if strings.Contains(err.Error(), "cannot unmarshal") || strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "password") {
		statusCode = http.StatusBadRequest
		errorDetail = ErrorDetail{
			Status: statusCode,
			Detail: err.Error(),
		}
	} else if strings.Contains(err.Error(), "1062") {
		statusCode = http.StatusConflict
		errorDetail = ErrorDetail{
			Status: statusCode,
			Detail: err.Error(),
		}
	} else if strings.Contains(err.Error(), "not found") {
		statusCode = http.StatusNotFound
		errorDetail = ErrorDetail{
			Status: statusCode,
			Detail: err.Error(),
		}
	} else if strings.Contains(err.Error(), "unauthorized") {
		statusCode = http.StatusUnauthorized
		errorDetail = ErrorDetail{
			Status: statusCode,
			Detail: err.Error(),
		}
	}

	errorResponse := ErrorResponse{
		Errors: []ErrorDetail{
			errorDetail,
		},
	}

	return errorResponse, statusCode
}
