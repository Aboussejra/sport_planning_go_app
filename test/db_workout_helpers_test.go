package test

import (
	"sport_planning_go_app/models"
	"testing"
)

func TestAddWorkout(t *testing.T) {
	db := setupTestDB(t)

	exercise1, _ := models.AddExercise(db, "Push Up")
	exercise2, _ := models.AddExercise(db, "Pull Up")

	workout, err := models.AddWorkout(db, "Full Body Workout", "Monday", []uint{exercise1.ID, exercise2.ID})
	if err != nil {
		t.Fatalf("models.AddWorkout failed: %v", err)
	}

	if workout.Name != "Full Body Workout" {
		t.Errorf("expected workout name to be 'Full Body Workout', got '%s'", workout.Name)
	}

	if len(workout.Exercises) != 2 {
		t.Errorf("expected 2 exercises in workout, got %d", len(workout.Exercises))
	}

	// Test uniqueness constraint
	_, err = models.AddWorkout(db, "Full Body Workout", "Tuesday", []uint{exercise1.ID})
	if err == nil {
		t.Errorf("expected models.AddWorkout to fail due to unique constraint, but it did not")
	}
}

func TestGetWorkoutExercises(t *testing.T) {
	db := setupTestDB(t)

	exercise1, _ := models.AddExercise(db, "Push Up")
	exercise2, _ := models.AddExercise(db, "Pull Up")

	workout, _ := models.AddWorkout(db, "Full Body Workout", "Monday", []uint{exercise1.ID, exercise2.ID})

	exercises, err := models.GetWorkoutExercises(db, workout.ID)
	if err != nil {
		t.Fatalf("GetWorkoutExercises failed: %v", err)
	}

	if len(exercises) != 2 {
		t.Errorf("expected 2 exercises, got %d", len(exercises))
	}
}
