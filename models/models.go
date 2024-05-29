package models

func AddWorkout(name string, reps int, sets int) error {
	_, err := db.Exec("INSERT INTO workouts (name, reps, sets) VALUES (?, ?, ?)", name, reps, sets)
	return err
}

// TODO: write test dummy DB and do real code
func ListAllWorkoutsDb() []Workout {
	// rows, err := db.Query("SELECT * FROM workouts")
	// log.Fatal(rows)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// An album slice to hold data from returned rows.
	// var workouts []Workout
	// for rows.Next() {
	// 	var workout Workout
	// 	if err := rows.Scan(&workout.Name); err != nil {
	// 		return workouts
	// 	}
	// }
	workouts := []Workout{
		{
			ID:   1,
			Name: "Leg Day",
			Exercises: []Exercise{
				{ID: 1, Name: "Squat"},
				{ID: 3, Name: "Lunges"},
			},
			DayOfWeek: "Monday",
		},
		{
			ID:   2,
			Name: "Chest Day",
			Exercises: []Exercise{
				{ID: 6, Name: "Chest Fly"},
			},
			DayOfWeek: "Wednesday",
		},
	}
	return workouts
}
