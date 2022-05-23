package rest_errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errorMessage = "error message"
)

func TestNewError(t *testing.T) {
	err := NewError(errorMessage)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorMessage, err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError(errorMessage)
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError(errorMessage)
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, "not_found", err.Error)
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError(errorMessage, NewError("database_error"))
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 1, len(err.Causes))
	assert.Equal(t, "database_error", err.Causes[0])
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError(errorMessage)
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Message)
	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, "unauthorized", err.Error)
}
