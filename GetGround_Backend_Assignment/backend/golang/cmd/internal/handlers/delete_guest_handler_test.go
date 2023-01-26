package handlers

import (
	"errors"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

func TestDeleteGuest_Success(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("DeleteGuest", mock.Anything, mock.Anything).Return(nil)

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("DELETE", "/guest/John", nil)
	w := httptest.NewRecorder()

	// add the variable to the request context
	r = mux.SetURLVars(r, map[string]string{
		"name": "John",
	})

	// call the handler
	handler.DeleteGuest(w, r)

	// assert the response
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
}

func TestDeleteGuest_Missing_Name(t *testing.T) {
	service := &mocks.GetGroundService{}

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("DELETE", "/guest/", nil)
	w := httptest.NewRecorder()

	// call the handler
	handler.DeleteGuest(w, r)

	// assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "missing user name. Name can not be empty", strings.TrimSuffix(w.Body.String(), "\n"))
}

func TestDeleteGuest_Failure(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("DeleteGuest", mock.Anything, mock.Anything).Return(errors.New("unable to delete guest"))

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("DELETE", "/guest/John", nil)
	w := httptest.NewRecorder()

	// add the variable to the request context
	r = mux.SetURLVars(r, map[string]string{
		"name": "John",
	})

	// call the handler
	handler.DeleteGuest(w, r)

	// assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "unable to delete guest", strings.TrimSuffix(w.Body.String(), "\n"))
}
