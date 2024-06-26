package handlers

import (
	"fmt"
	"net/http"
	"sport_planning_go_app/models"
	"sport_planning_go_app/views"
	"time"

	"github.com/a-h/templ"
)

// TODO Move to helper package
// errorHandlerWritesToHTTP is a helper function that checks for errors and writes an HTTP error response if an error is found.
func errorHandlerWritesToHTTP(w http.ResponseWriter, err error, msg string) bool {
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s: %v", msg, err))
		return true
	}
	return false
}
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
	exerciseName := r.Form.Get("exercise_name")
	workoutName := r.Form.Get("workout_name")

	// Get workout by name
	workout, err := models.GetWorkoutByName(models.DB, workoutName)
	if errorHandlerWritesToHTTP(w, err, "Workout not fetched successfully") {
		return
	}

	// Add exercise to workout
	exercise, err := models.AddExerciseToWorkout(models.DB, workout.ID, exerciseName)
	if errorHandlerWritesToHTTP(w, err, "Exercise not created successfully") {
		return
	}

	// Respond with success message
	fmt.Fprintf(w, "Exercise %s created successfully and added to workout %s", exercise.Name, workout.Name)
}

func WebPageAddExercisesToExistingWorkoutHandlers(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.AddExercisesToWorkout(models.ListAllWorkouts(models.DB), models.ListAllExercices(models.DB))).ServeHTTP(w, r)
}

func MainViewHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.MainView()).ServeHTTP(w, r)
}

func ViewExerciseInWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	workoutName := r.Form.Get("workout_name")

	workout, err := models.GetWorkoutByName(models.DB, workoutName)
	if errorHandlerWritesToHTTP(w, err, "Workout not fetched successfully") {
		return
	}
	templ.Handler(views.ViewExerciseInWorkout(workout.Exercises)).ServeHTTP(w, r)
}

func GetCurrentDayString(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "%s", currentTime.Weekday().String())
}

func ViewUndergoingWorkouts(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	workoutCurrentDay, err := models.GetWorkoutByDay(models.DB, currentDate.Weekday().String())
	if errorHandlerWritesToHTTP(w, err, fmt.Sprintf("Could not find undergoing workouts")) {
		return
	}
	workoutsToView := []models.Workout{workoutCurrentDay}
	templ.Handler(views.ViewUndergoingWorkouts(workoutsToView)).ServeHTTP(w, r)
}
