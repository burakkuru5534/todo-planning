package api

import (
	"encoding/json"
	"github.com/burakkuru5534/todo-planning/internal/model"
	"github.com/burakkuru5534/todo-planning/internal/util"
)

type Provider1 struct{}

func (p *Provider1) FetchTasks() ([]model.Task, error) {
	client := util.CreateHTTPClient()
	resp, err := client.Get("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
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
