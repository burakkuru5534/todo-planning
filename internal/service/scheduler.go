package service

import (
	"github.com/burakkuru5534/todo-planning/internal/model"
	"sort"
)

type Developer struct {
	Name          string
	HourlyRate    int
	MaxHours      int
	AssignedTasks []model.Task
}

type TaskScheduler struct {
	developers []Developer
	tasks      []model.Task
}

func NewTaskScheduler(developers []Developer, tasks []model.Task) *TaskScheduler {
	return &TaskScheduler{developers: developers, tasks: tasks}
}

func (ts *TaskScheduler) ScheduleTasks() (map[string][]model.Task, int) {
	// Sort tasks by difficulty (highest difficulty first)
	sort.Slice(ts.tasks, func(i, j int) bool {
		return ts.tasks[i].Difficulty > ts.tasks[j].Difficulty
	})

	devTasks := make(map[string][]model.Task)

	// Assign tasks to developers
	for _, task := range ts.tasks {
		assigned := false
		for i := range ts.developers {
			developer := &ts.developers[i]
			if developer.MaxHours >= task.DurationHours {
				developer.AssignedTasks = append(developer.AssignedTasks, task)
				developer.MaxHours -= task.DurationHours
				devTasks[developer.Name] = developer.AssignedTasks
				assigned = true
				break
			}
		}

		if !assigned {
			// If the task couldn't be assigned, append it to the developer with the highest available hours
			sort.Slice(ts.developers, func(i, j int) bool {
				return ts.developers[i].MaxHours > ts.developers[j].MaxHours
			})
			developer := &ts.developers[0]
			developer.AssignedTasks = append(developer.AssignedTasks, task)
			developer.MaxHours -= task.DurationHours
			devTasks[developer.Name] = developer.AssignedTasks
		}
	}

	// Calculate total weeks
	totalWeeks := 0
	for _, developer := range ts.developers {
		assignedHours := 45 - developer.MaxHours
		weeks := assignedHours / 45
		if assignedHours%45 > 0 {
			weeks++
		}
		if weeks > totalWeeks {
			totalWeeks = weeks
		}
	}

	return devTasks, totalWeeks
}
