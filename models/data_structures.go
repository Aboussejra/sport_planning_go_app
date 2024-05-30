package models

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

type Exercise struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Workout struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"unique;not null"`
	DayOfWeek string     `gorm:"not null"`
	Exercises []Exercise `gorm:"many2many:workout_exercises;"`
}

type WorkoutExercise struct {
	WorkoutID  uint
	ExerciseID uint
}

type ExerciseExecution struct {
	ID            uint     `gorm:"primaryKey"`
	ExerciseID    uint     `gorm:"not null"`
	WorkoutID     uint     `gorm:"not null"`
	Reps          int      `gorm:"not null"`
	ExecutionDate string   `gorm:"type:datetime;not null"`
	Exercise      Exercise `gorm:"foreignKey:ExerciseID"`
	Workout       Workout  `gorm:"foreignKey:WorkoutID"`
}

// WeeklyPlan represents the week's objectives with a set of workouts
type WeeklyPlan struct {
	Workouts map[DayOfWeek]Workout
}
