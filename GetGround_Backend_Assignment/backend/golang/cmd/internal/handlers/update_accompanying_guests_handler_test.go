package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateAccompanyingGuests_Success(t *testing.T) {
	amountOfGuests := 4
	jsonStr := fmt.Sprintf(`{"accompanying_guests": %d}`, amountOfGuests)
	jsonBytes := []byte(jsonStr)

	r, err := http.NewRequest("PUT", "/guests/John", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}
	r = mux.SetURLVars(r, map[string]string{"name": "John"})

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	service := &mocks.GetGroundService{}
	service.On("UpdateAccompanyingGuests", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create a new GuestHandler with the mock ground service
	handler := &GuestHandler{
		groundService: service,
	}

	// Call the UpdateAccompanyingGuests function
	handler.UpdateAccompanyingGuests(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Assert that the response body is a JSON object with a "name" key
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := response["name"]; !ok {
		t.Errorf("Expected response to have key 'name'")
	}

	// Assert that the value of the "name" key is "John"
	if response["name"] != "John" {
		t.Errorf("Expected name to be 'John', got %s", response["name"])
	}
}

func TestUpdateAccompanyingGuests_Missing_Name(t *testing.T) {
	amountOfGuests := 4
	jsonStr := fmt.Sprintf(`{"accompanying_guests": %d}`, amountOfGuests)
	jsonBytes := []byte(jsonStr)

	r, err := http.NewRequest("PUT", "/guests/", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	service := &mocks.GetGroundService{}
	service.On("UpdateAccompanyingGuests", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create a new GuestHandler with the mock ground service
	handler := &GuestHandler{
		groundService: service,
	}

	// Call the UpdateAccompanyingGuests function
	handler.UpdateAccompanyingGuests(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestUpdateAccompanyingGuests_Invalid_Accompanying_Amount(t *testing.T) {
	amountOfGuests := -3
	jsonStr := fmt.Sprintf(`{"accompanying_guests": %d}`, amountOfGuests)
	jsonBytes := []byte(jsonStr)

	r, err := http.NewRequest("PUT", "/guests/John", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}
	r = mux.SetURLVars(r, map[string]string{"name": "John"})

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	service := &mocks.GetGroundService{}
	service.On("UpdateAccompanyingGuests", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create a new GuestHandler with the mock ground service
	handler := &GuestHandler{
		groundService: service,
	}

	// Call the UpdateAccompanyingGuests function
	handler.UpdateAccompanyingGuests(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestUpdateAccompanyingGuests_Failure(t *testing.T) {
	amountOfGuests := 4
	jsonStr := fmt.Sprintf(`{"accompanying_guests": %d}`, amountOfGuests)
	jsonBytes := []byte(jsonStr)

	r, err := http.NewRequest("PUT", "/guests/John", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}
	r = mux.SetURLVars(r, map[string]string{"name": "John"})

	// Create a new response recorder to record the response
	w := httptest.NewRecorder()

	service := &mocks.GetGroundService{}
	service.On("UpdateAccompanyingGuests", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("forced for test"))

	// Create a new GuestHandler with the mock ground service
	handler := &GuestHandler{
		groundService: service,
	}

	// Call the UpdateAccompanyingGuests function
	handler.UpdateAccompanyingGuests(w, r)

	// Assert that the response has a status code of 200
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
