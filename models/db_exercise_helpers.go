package models

import (
	"gorm.io/gorm"
)

// AddExercise adds a new exercise
func AddExercise(db *gorm.DB, name string) (Exercise, error) {
	exercise := Exercise{Name: name}
	if err := db.Create(&exercise).Error; err != nil {
		return Exercise{}, err
	}
	return exercise, nil
}

// AddExerciseToWorkout adds an exercise to an existing workout
func AddExerciseToWorkout(db *gorm.DB, workoutID uint, name string) (Exercise, error) {
	var exercise Exercise

	if err := db.Where("name = ?", name).First(&exercise).Error; err != nil {
		return Exercise{}, err
	}

	// Add the exercise to the workout
	workoutExercise := WorkoutExercise{WorkoutID: workoutID, ExerciseID: exercise.ID}
	if err := db.Create(&workoutExercise).Error; err != nil {
		return Exercise{}, err
	}

	return exercise, nil
}

// AddExerciseExecution adds a new exercise execution to a workout exercise
func AddExerciseExecution(db *gorm.DB, exerciseID, workoutID uint, reps int, weight int, executionDate string) (ExerciseExecution, error) {
	execution := ExerciseExecution{
		ExerciseID:    exerciseID,
		WorkoutID:     workoutID,
		Reps:          reps,
		Weight:        weight,
		ExecutionDate: executionDate,
	}
	if err := db.Create(&execution).Error; err != nil {
		return ExerciseExecution{}, err
	}
	return execution, nil
}

// LoadExerciseExecutionsInWorkout loads all exercise executions for a given workout and exercise
func LoadExerciseExecutionsInWorkout(db *gorm.DB, workoutID, exerciseID uint) ([]ExerciseExecution, error) {
	var executions []ExerciseExecution
	if err := db.Where("workout_id = ? AND exercise_id = ?", workoutID, exerciseID).Find(&executions).Error; err != nil {
		return nil, err
	}
	return executions, nil
}

// LoadAllExerciseExecutions loads all exercise executions for a given exercise in all workouts
func LoadAllExerciseExecutions(db *gorm.DB, exerciseID uint) ([]ExerciseExecution, error) {
	var executions []ExerciseExecution
	if err := db.Where("exercise_id = ?", exerciseID).Find(&executions).Error; err != nil {
		return nil, err
	}
	return executions, nil
}
