package handlers

import (
	"encoding/json"
	"errors"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmptySeats_Success(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetEmptySeats", mock.Anything).Return(10, nil)

	// Create a new handler
	handler := &TableHandler{
		tableUseCase: service,
	}

	// Create a new HTTP request for the test
	r, err := http.NewRequest("GET", "/empty_seats", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	// Call the GetEmptySeats function
	handler.GetEmptySeats(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Assert that the response body is a JSON object with a "seats_empty" key
	var response map[string]int
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := response["seats_empty"]; !ok {
		t.Errorf("Expected response to have key 'seats_empty'")
	}

	// Assert that the value of the "seats_empty" key is 10
	if response["seats_empty"] != 10 {
		t.Errorf("Expected seats_empty to be 10, got %d", response["seats_empty"])
	}
}

func TestGetEmptySeats_Failure(t *testing.T) {
	service := &mocks.GetGroundService{}
	service.On("GetEmptySeats", mock.Anything).Return(0, errors.New("forced for test"))

	// Create a new handler
	handler := &TableHandler{
		tableUseCase: service,
	}

	// Create a new HTTP request for the test
	r, err := http.NewRequest("GET", "/empty_seats", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	// Call the GetEmptySeats function
	handler.GetEmptySeats(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
