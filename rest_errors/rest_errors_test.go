package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TODO: Complete errors test

// Old test
// func TestNewError(t *testing.T) {
// 	error := errors.New("this is a message")
// 	assert.NotNil(t, error)
// 	assert.EqualValues(t, "this is a message", error.Error())
// }

func TestInternalRequestError(t *testing.T) {
	err := NewInternalServerError("this is a message", errors.New("database example error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 500 - error: database example error - causes: [ [database example error] ]", err.Error())
	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database example error", err.Causes()[0])

	// errBytes, _ := json.Marshal(err)
	// fmt.Println(string(errBytes))
}
func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a message", errors.New("bad request error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 400 - error: bad request error - causes: [ [] ]", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a message", errors.New("Not found error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 404 - error: Not found error - causes: [ [] ]", err.Error())
}
