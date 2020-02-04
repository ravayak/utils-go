package rest_errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	error := NewError("this is a message")
	assert.NotNil(t, error)
	assert.EqualValues(t, "this is a message", error.Error())
}

func TestInternalRequestError(t *testing.T) {
	err := NewInternalServerError("this is a message", NewError("database example error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database example error", err.Causes[0])

	// errBytes, _ := json.Marshal(err)
	// fmt.Println(string(errBytes))
}
func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a message", NewError("bad request error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "bad request error", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a message", NewError("Not found error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "Not found error", err.Error)
}
