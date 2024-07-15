package database

import (
	"github.com/burakkuru5534/todo-planning/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.Task{})
}

func GetDB() *gorm.DB {
	return db
}
