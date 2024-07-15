package service

import (
	"github.com/burakkuru5534/todo-planning/internal/api"
	"github.com/burakkuru5534/todo-planning/internal/model"
	"sync"

	"github.com/jinzhu/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (ts *TaskService) FetchAndStoreTasks() {
	var wg sync.WaitGroup
	taskChan := make(chan []model.Task)

	providers := []interface {
		FetchTasks() ([]model.Task, error)
	}{
		&api.Provider1{},
		&api.Provider2{},
	}

	for _, provider := range providers {
		wg.Add(1)
		go func(p interface{ FetchTasks() ([]model.Task, error) }) {
			defer wg.Done()
			tasks, err := p.FetchTasks()
			if err != nil {
				return
			}
			taskChan <- tasks
		}(provider)
	}

	go func() {
		wg.Wait()
		close(taskChan)
	}()

	for tasks := range taskChan {
		for _, task := range tasks {
			ts.db.Create(&task)
		}
	}
}

func (ts *TaskService) GetTasks() []model.Task {
	var tasks []model.Task
	ts.db.Find(&tasks)
	return tasks
}
