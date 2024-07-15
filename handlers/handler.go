package handlers

import (
	"encoding/json"
	"github.com/burakkuru5534/todo-planning/services"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func FetchTasks(w http.ResponseWriter, r *http.Request) {
	err := services.FetchAndStoreTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
