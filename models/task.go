package models

type Task struct {
	Name       string `json:"name"`
	Duration   int    `json:"duration"`
	Difficulty int    `json:"difficulty"`
}
