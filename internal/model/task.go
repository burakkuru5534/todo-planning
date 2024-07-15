package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name       string `json:"name"`
	Duration   int    `json:"duration"`
	Difficulty int    `json:"difficulty"`
}
