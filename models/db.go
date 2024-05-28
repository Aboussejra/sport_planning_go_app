package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDb() {
	var err error
	db, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDb() {
	db.Close()
}
func SaveWorkout(workout *Workout) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert the workout
	result, err := tx.Exec("INSERT INTO workouts (name, day_of_week) VALUES (?, ?)", workout.Name, workout.DayOfWeek)
	if err != nil {
		return err
	}
	workoutID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Insert the exercises
	for _, exercise := range workout.Exercises {
		result, err = tx.Exec("INSERT INTO exercises (name) VALUES (?)", exercise.Name)
		if err != nil {
			return err
		}
		exerciseID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// Associate the exercise with the workout
		_, err = tx.Exec("INSERT INTO workout_exercises (workout_id, exercise_id) VALUES (?, ?)", workoutID, exerciseID)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
