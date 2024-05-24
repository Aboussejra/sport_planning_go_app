package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl.Execute(w, nil)
}

func createWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract form values
	name := r.Form.Get("name")
	dayOfWeek := r.Form.Get("dayOfWeek")

	// Validate form data
	if name == "" || dayOfWeek == "" {
		http.Error(w, "Name and Day of the Week are required fields", http.StatusBadRequest)
		return
	}

	// Create the workout object
	workout := Workout{
		Name:      name,
		DayOfWeek: dayOfWeek,
	}
	// Implement database logic here
	saveWorkout(db, &workout)
	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Workout created successfully: %s on %s with %d exercises", workout.Name, workout.DayOfWeek, len(workout.Exercises))
}
