package handlers

import (
	"bytes"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddGuestToGuestList_Valid_Data(t *testing.T) {
	// Initialize test data
	name := "John Doe"
	table := 1
	guests := 2
	jsonStr := fmt.Sprintf(`{"table": %d, "accompanying_guests": %d}`, table, guests)
	jsonBytes := []byte(jsonStr)

	// Create a new request
	req, err := http.NewRequest("POST", "/guestlist/"+name, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	service := &mocks.GetGroundService{}
	service.On("AddGuestToList", mock.Anything, mock.Anything).Return(nil)

	// Create a new handler
	handler := &GuestHandler{
		groundService: service,
	}

	// Create a new mux router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/guestlist/{name}", handler.AddGuestToGuestList).Methods("POST")

	// Execute the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := fmt.Sprintf(`{"name":"%s"}`, name)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAddGuestToGuestList_Invalid_tableId(t *testing.T) {
	// Initialize test data
	name := "John Doe"
	table := 0
	guests := 2
	jsonStr := fmt.Sprintf(`{"table": %d, "accompanying_guests": %d}`, table, guests)
	jsonBytes := []byte(jsonStr)

	// Create a new request
	req, err := http.NewRequest("POST", "/guestlist/"+name, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new handler
	handler := &GuestHandler{
		groundService: &mocks.GetGroundService{},
	}

	// Create a new mux router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/guestlist/{name}", handler.AddGuestToGuestList).Methods("POST")

	// Execute the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestAddGuestToGuestList_Invalid_accompanying_guests(t *testing.T) {
	// Initialize test data
	name := "John Doe"
	table := 1
	guests := -3
	jsonStr := fmt.Sprintf(`{"table": %d, "accompanying_guests": %d}`, table, guests)
	jsonBytes := []byte(jsonStr)

	// Create a new request
	req, err := http.NewRequest("POST", "/guestlist/"+name, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new handler
	handler := &GuestHandler{
		groundService: &mocks.GetGroundService{},
	}

	// Create a new mux router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/guestlist/{name}", handler.AddGuestToGuestList).Methods("POST")

	// Execute the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
