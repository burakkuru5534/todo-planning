package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	ID            int    `json:"id"`
	Name          string `json:"name"`
	DurationHours int    `json:"duration_hours"`
	Difficulty    int    `json:"difficulty"`
}
