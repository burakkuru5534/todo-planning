package service

import (
	"github.com/burakkuru5534/todo-planning/internal/model"
	"testing"
)

func TestTaskScheduler_ScheduleTasks(t *testing.T) {
	developers := []Developer{
		{Name: "DEV1", HourlyRate: 1, MaxHours: 45},
		{Name: "DEV2", HourlyRate: 2, MaxHours: 45},
		{Name: "DEV3", HourlyRate: 3, MaxHours: 45},
		{Name: "DEV4", HourlyRate: 4, MaxHours: 45},
		{Name: "DEV5", HourlyRate: 5, MaxHours: 45},
	}

	tasks := []model.Task{
		{Name: "Task1", DurationHours: 5, Difficulty: 3},
		{Name: "Task2", DurationHours: 10, Difficulty: 1},
		{Name: "Task3", DurationHours: 15, Difficulty: 2},
		{Name: "Task4", DurationHours: 20, Difficulty: 5},
		{Name: "Task5", DurationHours: 25, Difficulty: 4},
	}

	ts := NewTaskScheduler(developers, tasks)
	devTasks, totalWeeks := ts.ScheduleTasks()

	// Test total number of developers
	if len(devTasks) != 5 {
		t.Errorf("expected 5 developers, got %d", len(devTasks))
	}

	// Test total number of weeks
	expectedTotalWeeks := 2 // Assuming this value based on the provided tasks and developers
	if totalWeeks != expectedTotalWeeks {
		t.Errorf("expected total weeks %d, got %d", expectedTotalWeeks, totalWeeks)
	}

	// Test task distribution per developer
	expectedTaskDistribution := map[string]int{
		"DEV1": 1,
		"DEV2": 1,
		"DEV3": 0,
		"DEV4": 1,
		"DEV5": 2,
	}

	for dev, tasks := range devTasks {
		if len(tasks) != expectedTaskDistribution[dev] {
			t.Errorf("expected %d tasks for %s, got %d", expectedTaskDistribution[dev], dev, len(tasks))
		}
	}
}
