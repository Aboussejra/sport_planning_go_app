package models

import "log"

func AddWorkout(name string, reps int, sets int) error {
	_, err := db.Exec("INSERT INTO workouts (name, reps, sets) VALUES (?, ?, ?)", name, reps, sets)
	return err
}

// TODO: write test dummy DB and do real code, Migrate GORM?
func ListAllWorkoutsDb() []Workout {
	// Query to select all workouts
	rows, err := db.Query("SELECT name FROM workouts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to hold the workouts
	var workouts []Workout

	// Iterate over the rows
	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.Name); err != nil {
			log.Fatal(err)
		}

		var exercises []Exercise
		workout.Exercises = exercises
		workouts = append(workouts, workout)
	}
	return workouts
}
