package handlers

import (
	"encoding/json"
	"errors"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGuestListHandler_Success(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetGuestsList", mock.Anything).Return([]models.Guest{{Name: "John", Table: 1, AccompanyingGuests: 2}}, nil, nil)

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("GET", "/guest_list", nil)
	w := httptest.NewRecorder()

	// call the handler
	handler.GetGuestsList(w, r)

	// assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	var guests []models.Guest
	json.Unmarshal(w.Body.Bytes(), &guests)
	assert.Equal(t, []models.Guest{{Name: "John", Table: 1, AccompanyingGuests: 2}}, guests)
}

func TestGetGuestListHandler_Failure(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetGuestsList", mock.Anything).Return([]models.Guest{}, errors.New("forced for test"))

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// create request and response recorder
	r, _ := http.NewRequest("GET", "/guest_list", nil)
	w := httptest.NewRecorder()

	// call the handler
	handler.GetGuestsList(w, r)

	// assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
