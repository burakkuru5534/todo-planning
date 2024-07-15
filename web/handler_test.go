package handler_test

import (
	"bytes"
	"github.com/burakkuru5534/todo-planning/internal/service"
	handler "github.com/burakkuru5534/todo-planning/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/burakkuru5534/todo-planning/internal/model"
)

// MockTaskService is a mock implementation of service.TaskService interface
type MockTaskService struct{}

func (mts *MockTaskService) GetTasks() []model.Task {
	// Mocked tasks
	tasks := []model.Task{
		{ID: 1, Name: "Task 1", DurationHours: 3, Difficulty: 1},
		{ID: 2, Name: "Task 2", DurationHours: 2, Difficulty: 2},
		{ID: 3, Name: "Task 3", DurationHours: 5, Difficulty: 3},
	}
	return tasks
}

func TestIndexHandler(t *testing.T) {
	// Mock TaskService
	mockTaskService := &service.TaskService{}
	mockTaskService.GetTasks()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Create an IndexHandler instance with the mock TaskService
	handlerFunc := handler.IndexHandler(mockTaskService)

	// Serve the HTTP request to the handler
	handlerFunc(rec, req)

	// Check the status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Check the response body (assuming it's HTML, we can check for certain strings)
	expectedSubstring := "<h1>Developer Task Assignments</h1>"
	if !bytes.Contains(rec.Body.Bytes(), []byte(expectedSubstring)) {
		t.Errorf("Expected response body to contain %q, got:\n%s", expectedSubstring, rec.Body.String())
	}
}
