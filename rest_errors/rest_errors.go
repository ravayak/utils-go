package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type restError struct {
	message string        `json:"message"`
	status  int           `json:"code"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

// func (e restError) Error() string {
// 	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
// 		e.message, e.status, e.error, e.causes)
// }

func (e restError) Error() string {
	return e.error
}

func (e restError) Message() string {
	return e.message
}

func (e restError) Status() int {
	return e.status
}

func (e restError) NewError(msg string) error {
	return errors.New(msg)
}

func (e restError) Causes() []interface{} {
	return e.causes
}

func NewRestError(message string, status int, err string, causes []interface{}) RestError {
	return restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   err,
		causes:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestError, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string, err error) RestError {
	return restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   err.Error(),
	}
}

func NewNotFoundError(message string, err error) RestError {
	return restError{
		message: message,
		status:  http.StatusNotFound,
		error:   err.Error(),
	}
}

func NewUnauthorizedError(message string, err error) RestError {
	return restError{
		message: message,
		status:  http.StatusUnauthorized,
		error:   err.Error(),
	}
}

func NewInternalServerError(message string, err error) RestError {
	result := restError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   err.Error(),
	}

	if err != nil {
		result.causes = append(result.causes, err.Error())
	}

	return result

}
