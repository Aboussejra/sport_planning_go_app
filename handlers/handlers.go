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
func SaveWorkout(w http.ResponseWriter, r *http.Request) {
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
	workout := models.Workout{
		Name:      name,
		DayOfWeek: dayOfWeek,
	}
	// Implement database logic here
	models.SaveWorkout(&workout)
	// Respond with success message
	fmt.Fprintf(w, "Workout created successfully: %s on %s with %d exercises", workout.Name, workout.DayOfWeek, len(workout.Exercises))
}

func AddExercisesToWorkoutHandlers(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.AddExercises(models.ListAllWorkoutsDb())).ServeHTTP(w, r)
}
