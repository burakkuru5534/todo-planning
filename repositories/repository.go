package repositories

import (
	"github.com/burakkuru5534/todo-planning/database"
	"github.com/burakkuru5534/todo-planning/models"
)

func SaveTasks(tasks []models.Task) error {
	db := database.GetDB()
	for _, task := range tasks {
		if err := db.Create(&task).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	db := database.GetDB()
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
