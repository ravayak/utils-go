package rest_errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Message string        `json:"message"`
	Status  int           `json:"code"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestError(message string, err error) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   err.Error(),
	}
}

func NewNotFoundError(message string, err error) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   err.Error(),
	}
}

func NewInternalServerError(message string, err error) *RestError {
	result := &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result

}

func NewError(msg string) error {
	return errors.New(msg)
}
