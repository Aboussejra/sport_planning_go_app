package main

import (
	"log"
	"net/http"
	"sport_planning_go_app/handlers"
	"sport_planning_go_app/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.InitDB()
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/create-workout", handlers.CreateWorkoutHandler)
	http.HandleFunc("/save-workout", handlers.SaveWorkoutHandler)
	http.HandleFunc("/retrieve-html-to-add-exercises-to-existing-workouts", handlers.WebPageAddExercisesToExistingWorkoutHandlers)
	http.HandleFunc("/add-exercise", handlers.AddExerciseHandler)
	http.HandleFunc("/view-exercises-in-workout", handlers.ViewExerciseInWorkoutHandler)
	http.HandleFunc("/home-view", handlers.MainViewHandler)
	http.HandleFunc("/view-workout-identified-by-day", handlers.ViewWorkoutGivenDay)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
