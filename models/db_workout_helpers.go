package models

import "gorm.io/gorm"

// GetWorkoutExercises retrieves exercises for a given workout
func GetWorkoutExercises(db *gorm.DB, workoutID uint) ([]Exercise, error) {
	var workout Workout
	if err := db.Preload("Exercises").First(&workout, workoutID).Error; err != nil {
		return nil, err
	}
	return workout.Exercises, nil
}

// AddWorkout adds a new workout with a list of exercises
func AddWorkout(db *gorm.DB, name string, dayOfWeek string, exerciseIDs []uint) (Workout, error) {
	var exercises []Exercise
	if err := db.Find(&exercises, exerciseIDs).Error; err != nil {
		return Workout{}, err
	}

	workout := Workout{
		Name:      name,
		DayOfWeek: dayOfWeek,
		Exercises: exercises,
	}

	if err := db.Create(&workout).Error; err != nil {
		return Workout{}, err
	}

	return workout, nil
}

// ListAllWorkoutsDb lists all workouts from the database
func ListAllWorkouts(db *gorm.DB) ([]Workout, error) {
	var workouts []Workout
	db.Find(&workouts)
	return workouts, nil
}
