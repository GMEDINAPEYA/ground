package handlers

import (
	"encoding/json"
	"errors"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArrivedGuests(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetArrivedGuests", mock.Anything).Return([]models.ArrivedGuest{{Name: "John", TimeArrived: "15:00"}}, nil, nil)

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("GET", "/guest/arrived", nil)
	w := httptest.NewRecorder()

	// call the handler
	handler.GetArrivedGuests(w, r)

	// assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	var guests []models.ArrivedGuest
	json.Unmarshal(w.Body.Bytes(), &guests)
	assert.Equal(t, []models.ArrivedGuest{{Name: "John", TimeArrived: "15:00"}}, guests)
}

func TestGetArrivedGuests_Error(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetArrivedGuests", mock.Anything).Return([]models.ArrivedGuest{}, errors.New("unable to fetch guests"))

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("GET", "/guest/arrived", nil)
	w := httptest.NewRecorder()

	// call the handler
	handler.GetArrivedGuests(w, r)

	// assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "unable to fetch guests", strings.TrimSuffix(w.Body.String(), "\n"))
}
