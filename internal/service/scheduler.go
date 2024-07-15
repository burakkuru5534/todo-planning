package service

import (
	"github.com/burakkuru5534/todo-planning/internal/model"
	"github.com/burakkuru5534/todo-planning/internal/util"
)

type Developer struct {
	Name       string // Developer'ın adı
	HourlyRate int    // Developer'ın saatlik iş birimi
	MaxHours   int    // Developer'ın haftalık maksimum çalışma saati
}

type TaskScheduler struct {
	Developers []Developer     // Developer listesi
	Tasks      []model.Task    // Görev listesi
	TaskCh     chan model.Task // Görevleri almak için channel
	CompleteCh chan bool       // Görev tamamlandığında işaretlemek için channel
}

func NewTaskScheduler(developers []Developer, tasks []model.Task) *TaskScheduler {
	return &TaskScheduler{
		Developers: developers,
		Tasks:      tasks,
		TaskCh:     make(chan model.Task),
		CompleteCh: make(chan bool),
	}
}

func (ts *TaskScheduler) ScheduleTasks() (map[string][]model.Task, int) {
	// Developer'lara görevlerin atanması
	devTasks := make(map[string][]model.Task)
	totalWeeks := 0

	// Görevleri zorluk derecesine göre sıralanması
	sortTasksByDifficulty(ts.Tasks)
	util.InfoLogger.Println("sortTasksByDifficulty", ts.Tasks)

	// Görevleri işlemek için goroutine'ler başlatılması
	for i := 0; i < len(ts.Tasks); i++ {
		go ts.processTask()
	}

	// Görevleri channel'a gönder
	for _, task := range ts.Tasks {
		ts.TaskCh <- task
	}

	// Görevlerin tamamlanmasını bekle
	for i := 0; i < len(ts.Tasks); i++ {
		<-ts.CompleteCh
	}

	// Haftalık iş bölümü yapma
	for _, task := range ts.Tasks {
		devName := task.AssignedTo
		devTasks[devName] = append(devTasks[devName], task)
	}

	// Toplam hafta hesaplama
	for _, dev := range ts.Developers {
		totalWeeks += (ts.getTotalAssignedHours(dev.Name) + dev.MaxHours - 1) / dev.MaxHours
	}

	return devTasks, totalWeeks
}

func (ts *TaskScheduler) processTask() {
	for {
		task, more := <-ts.TaskCh
		if more {
			for _, dev := range ts.Developers {
				if ts.canAssignTaskToDeveloper(dev, task) {
					task.AssignedTo = dev.Name
					break
				}
			}
			ts.CompleteCh <- true
		} else {
			return
		}
	}
}

func sortTasksByDifficulty(tasks []model.Task) []model.Task {
	if len(tasks) <= 1 {
		return tasks
	}

	// Choose the last element as pivot
	pivot := tasks[len(tasks)-1].Difficulty

	// Partition the array around the pivot
	left, right := 0, len(tasks)-1
	for i := 0; i < right; i++ {
		if tasks[i].Difficulty < pivot {
			tasks[left], tasks[i] = tasks[i], tasks[left]
			left++
		}
	}
	tasks[left], tasks[right] = tasks[right], tasks[left]

	// Recursively sort the partitions
	sortTasksByDifficulty(tasks[:left])
	sortTasksByDifficulty(tasks[left+1:])

	return tasks
}

func (ts *TaskScheduler) canAssignTaskToDeveloper(dev Developer, task model.Task) bool {
	// Developer'ın haftalık maksimum çalışma saatini aşmamış olmalı
	if ts.getTotalAssignedHours(dev.Name)+task.DurationHours > dev.MaxHours {
		return false
	}
	return true
}

func (ts *TaskScheduler) getTotalAssignedHours(devName string) int {
	totalHours := 0
	for _, task := range ts.Tasks {
		if task.AssignedTo == devName {
			totalHours += task.DurationHours
		}
	}
	return totalHours
}
