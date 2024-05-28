package models

import "time"

// DayOfWeek represents the day of the week
type DayOfWeek int

const (
	Monday DayOfWeek = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Exercise represents an exercise in the database
type Exercise struct {
	ID   int
	Name string
}

// ExerciseExecution represents the execution of an exercise with reps and execution date
type ExerciseExecution struct {
	ID            int
	ExerciseID    int
	Reps          int
	ExecutionDate time.Time
}

// Workout represents a workout, which is a collection of exercises that should be done in a given DayOfWeek
type Workout struct {
	ID        int
	Name      string
	Exercises []Exercise
	DayOfWeek string
}

// WeeklyPlan represents the week's objectives with a set of workouts
type WeeklyPlan struct {
	Workouts map[DayOfWeek]Workout
}
