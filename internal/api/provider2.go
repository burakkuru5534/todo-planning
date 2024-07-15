package api

import (
	"encoding/json"
	"github.com/burakkuru5534/todo-planning/internal/model"
	"github.com/burakkuru5534/todo-planning/internal/util"
)

type Provider2 struct{}

func (p *Provider2) FetchTasks() ([]model.Task, error) {
	client := util.CreateHTTPClient()
	resp, err := client.Get("https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []model.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
