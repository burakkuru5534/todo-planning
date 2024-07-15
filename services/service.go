package services

import (
	"encoding/json"
	"github.com/burakkuru5534/todo-planning/models"
	"github.com/burakkuru5534/todo-planning/repositories"
	"io/ioutil"
	"net/http"
)

var providers = []string{
	"https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd",
	"https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143",
}

func FetchAndStoreTasks() error {
	var allTasks []models.Task
	for _, provider := range providers {
		resp, err := http.Get(provider)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var tasks []models.Task
		if err := json.Unmarshal(body, &tasks); err != nil {
			return err
		}
		allTasks = append(allTasks, tasks...)
	}
	return repositories.SaveTasks(allTasks)
}

func GetTasks() ([]models.Task, error) {
	return repositories.GetTasks()
}
