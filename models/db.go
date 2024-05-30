package models

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("./database/database.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&Exercise{}, &ExerciseExecution{}, &Workout{})
}
