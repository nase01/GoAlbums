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

	// DB custom error message
	if strings.Contains(err.Error(), "1062") {
		statusCode = http.StatusConflict
		detail = "Duplicate entry"
	} else if strings.Contains(err.Error(), "not found") {
		statusCode = http.StatusNotFound
		detail = "Record(s) not found"
	}

	errorResponse := ErrorResponse{
		Errors: []ErrorDetail{
			{
				Status: statusCode,
				Detail: detail,
			},
		},
	}

	return errorResponse, statusCode
}
