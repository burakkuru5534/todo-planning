package main

import (
	"github.com/burakkuru5534/todo-planning/internal/model"
	"github.com/burakkuru5534/todo-planning/internal/service"
	"github.com/burakkuru5534/todo-planning/internal/util"
	handler "github.com/burakkuru5534/todo-planning/web"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Initialize database
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		util.ErrorLogger.Println("failed to connect to database:", err)
		return
	}
	defer db.Close()

	util.InfoLogger.Println("Gorm Open worked.")

	// Migrate the schema
	db.AutoMigrate(&model.Task{})
	util.InfoLogger.Println("AutoMigrate worked.")

	// Fetch tasks and store in database
	taskService := service.NewTaskService(db)
	util.InfoLogger.Println("NewTaskService fetched.")
	taskService.FetchAndStoreTasks()
	util.InfoLogger.Println("NewTaskService fetched and stored.")

	// Setup HTTP server
	http.HandleFunc("/", handler.IndexHandler(taskService))
	util.InfoLogger.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		util.ErrorLogger.Println("failed to start server:", err)
	}
}
