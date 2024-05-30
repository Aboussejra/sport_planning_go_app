package models

import (
	"testing"
)

func TestAddExercise(t *testing.T) {
	db := setupTestDB(t)

	exercise, err := AddExercise(db, "Push Up")
	if err != nil {
		t.Fatalf("AddExercise failed: %v", err)
	}

	if exercise.Name != "Push Up" {
		t.Errorf("expected exercise name to be 'Push Up', got '%s'", exercise.Name)
	}

	// Test uniqueness constraint
	_, err = AddExercise(db, "Push Up")
	if err == nil {
		t.Errorf("expected AddExercise to fail due to unique constraint, but it did not")
	}
}

func TestAddExerciseExecution(t *testing.T) {
	db := setupTestDB(t)

	exercise, _ := AddExercise(db, "Push Up")
	workout, _ := AddWorkout(db, "Full Body Workout", "Monday", []uint{exercise.ID})

	execution, err := AddExerciseExecution(db, exercise.ID, workout.ID, 15, "2024-05-30 10:00:00")
	if err != nil {
		t.Fatalf("AddExerciseExecution failed: %v", err)
	}

	if execution.Reps != 15 {
		t.Errorf("expected reps to be 15, got %d", execution.Reps)
	}
}

func TestLoadExerciseExecutionsInWorkout(t *testing.T) {
	db := setupTestDB(t)

	exercise, _ := AddExercise(db, "Push Up")
	workout, _ := AddWorkout(db, "Full Body Workout", "Monday", []uint{exercise.ID})

	AddExerciseExecution(db, exercise.ID, workout.ID, 15, "2024-05-30 10:00:00")

	executions, err := LoadExerciseExecutionsInWorkout(db, workout.ID, exercise.ID)
	if err != nil {
		t.Fatalf("LoadExerciseExecutionsInWorkout failed: %v", err)
	}

	if len(executions) != 1 {
		t.Errorf("expected 1 execution, got %d", len(executions))
	}

	if executions[0].Reps != 15 {
		t.Errorf("expected reps to be 15, got %d", executions[0].Reps)
	}
}

func TestLoadAllExerciseExecutions(t *testing.T) {
	db := setupTestDB(t)

	exercise, _ := AddExercise(db, "Push Up")
	workout1, _ := AddWorkout(db, "Full Body Workout", "Monday", []uint{exercise.ID})
	workout2, _ := AddWorkout(db, "Upper Body Workout", "Wednesday", []uint{exercise.ID})

	AddExerciseExecution(db, exercise.ID, workout1.ID, 15, "2024-05-30 10:00:00")
	AddExerciseExecution(db, exercise.ID, workout2.ID, 20, "2024-05-31 10:00:00")

	executions, err := LoadAllExerciseExecutions(db, exercise.ID)
	if err != nil {
		t.Fatalf("LoadAllExerciseExecutions failed: %v", err)
	}

	if len(executions) != 2 {
		t.Errorf("expected 2 executions, got %d", len(executions))
	}
}
