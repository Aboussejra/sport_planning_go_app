package main

import (
	"log"
	"net/http"
	"sport_planning_go_app/handlers"
	"sport_planning_go_app/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.InitDb()
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/create-workout", handlers.CreateWorkoutHandler)
	http.HandleFunc("/save-workout", handlers.SaveWorkout)
	http.HandleFunc("/add-exercices-to-existing-workouts", handlers.AddExercisesToWorkoutHandlers)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	models.CloseDb()
}
