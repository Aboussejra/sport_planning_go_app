package models

import "database/sql"

func addWorkout(db *sql.DB, name string, reps int, sets int) error {
	_, err := db.Exec("INSERT INTO workouts (name, reps, sets) VALUES (?, ?, ?)", name, reps, sets)
	return err
}
