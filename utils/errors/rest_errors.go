package errors

import "net/http"

type RestError struct {
	Message string `json:message`
	Status  int    `json:status`
	Error   string `json:error`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request!",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
