package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/burakkuru5534/todo-planning/internal/model"
)

func TestProvider1_FetchTasks(t *testing.T) {
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response JSON
		mockResponse := `[{"id": 1, "name": "Task 1", "duration_hours": 3, "difficulty": 1"}]`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Override the HTTP client with the mock server URL
	provider := Provider1{}

	// Call FetchTasks method
	tasks, err := provider.FetchTasks()

	// Check for errors
	if err != nil {
		t.Fatalf("FetchTasks returned error: %v", err)
	}

	// Check if tasks are returned correctly
	expectedTasks := []model.Task{{ID: 1, Name: "Task 1", Duration: 3, Difficulty: 1}}
	if len(tasks) != len(expectedTasks) {
		t.Fatalf("Expected %d tasks, got %d", len(expectedTasks), len(tasks))
	}
	for i := range expectedTasks {
		if tasks[i].ID != expectedTasks[i].ID || tasks[i].Name != expectedTasks[i].Name ||
			tasks[i].Duration != expectedTasks[i].Duration || tasks[i].Difficulty != expectedTasks[i].Difficulty {
			t.Errorf("Task %d mismatch:\nExpected: %+v\nGot: %+v", i+1, expectedTasks[i], tasks[i])
		}
	}
}
