package handler

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("http code %d, message: %s", e.Code, e.Message)
}

func NewInternalServerError() ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Errror",
	}
}
