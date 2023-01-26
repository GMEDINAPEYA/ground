package handlers

import (
	"bytes"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTable_Valid_Data(t *testing.T) {
	// Initialize test data
	capacity := 4
	jsonStr := fmt.Sprintf(`{"capacity": %d}`, capacity)
	jsonBytes := []byte(jsonStr)

	// Create a new request
	req, err := http.NewRequest("POST", "/table", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	table := &models.Table{
		Id:       1,
		Capacity: capacity,
	}

	service := &mocks.GetGroundService{}
	service.On("CreateTable", mock.Anything, mock.Anything).Return(table, nil)

	// Create a new handler
	handler := &TableHandler{
		tableUseCase: service,
	}

	// Create a new mux router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/table", handler.CreateTable).Methods("POST")

	// Execute the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := fmt.Sprintf(`{"id":1,"capacity":%d}`, capacity)
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateTable_Invalid_Capacity(t *testing.T) {
	// Initialize test data
	capacity := 0
	jsonStr := fmt.Sprintf(`{"capacity": %d}`, capacity)
	jsonBytes := []byte(jsonStr)

	// Create a new request
	req, err := http.NewRequest("POST", "/table", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	table := &models.Table{
		Id:       1,
		Capacity: capacity,
	}

	service := &mocks.GetGroundService{}
	service.On("CreateTable", mock.Anything, mock.Anything).Return(table, nil)

	// Create a new handler
	handler := &TableHandler{
		tableUseCase: service,
	}

	// Create a new mux router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/table", handler.CreateTable).Methods("POST")

	// Execute the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
