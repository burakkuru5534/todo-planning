package handler

import (
	"github.com/burakkuru5534/todo-planning/internal/model"
	"github.com/burakkuru5534/todo-planning/internal/service"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("web/static/index.html"))

func IndexHandler(ts *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := ts.GetTasks()

		// Mocked developers
		developers := []service.Developer{
			{Name: "DEV1", HourlyRate: 1, MaxHours: 45},
			{Name: "DEV2", HourlyRate: 2, MaxHours: 45},
			{Name: "DEV3", HourlyRate: 3, MaxHours: 45},
			{Name: "DEV4", HourlyRate: 4, MaxHours: 45},
			{Name: "DEV5", HourlyRate: 5, MaxHours: 45},
		}

		scheduler := service.NewTaskScheduler(developers, tasks)
		devTasks, totalWeeks := scheduler.ScheduleTasks()

		data := struct {
			DevTasks   map[string][]model.Task
			TotalWeeks int
		}{
			DevTasks:   devTasks,
			TotalWeeks: totalWeeks,
		}

		tmpl.Execute(w, data)
	}
}
