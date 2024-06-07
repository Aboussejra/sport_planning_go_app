package handlers

import (
	"fmt"
	"net/http"
	"sport_planning_go_app/models"
	"sport_planning_go_app/views"
	"time"

	"github.com/a-h/templ"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	templ.Handler(views.Index(time.Now())).ServeHTTP(w, r)
}

func CreateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.CreateWorkout()).ServeHTTP(w, r)
}
func SaveWorkoutHandler(w http.ResponseWriter, r *http.Request) {
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

	// Implement database logic here
	workout, err := models.AddWorkout(models.DB, name, dayOfWeek, []uint{})
	if err != nil {
		// Respond with success message
		fmt.Fprintf(w, "Workout not created succesfully due to %s", err)
	} else {
		// Respond with success message
		fmt.Fprintf(w, "Workout created successfully: %s on %s with %d exercises", workout.Name, workout.DayOfWeek, len(workout.Exercises))
	}
}
func AddExerciseHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	exercise_name := r.Form.Get("exercise_name")
	workout_name := r.Form.Get("workout_name")
	workout, err := models.GetWorkoutByName(models.DB, workout_name)
	exercise, err := models.AddExerciseToWorkout(models.DB, workout.ID, exercise_name)
	if err != nil {
		// Respond with success message
		fmt.Fprintf(w, "Exercise not created succesfully due to %s", err)
	} else {
		// Respond with success message
		fmt.Fprintf(w, "Exercise %s created successfully and added to workout %s", exercise.Name, workout.Name)
	}
}

func WebPageAddExercisesToExistingWorkoutHandlers(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.AddExercisesToWorkout(models.ListAllWorkouts(models.DB), models.ListAllExercices(models.DB))).ServeHTTP(w, r)
}

func CalendarMainViewHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.CalendarView()).ServeHTTP(w, r)
}
