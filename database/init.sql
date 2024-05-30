-- Create the Exercise table
CREATE TABLE exercises (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE
);

-- Create the Workout table
CREATE TABLE workouts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE,
    day_of_week VARCHAR(128) NOT NULL
);

-- Create the WorkoutExercises join table
CREATE TABLE workout_exercises (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    workout_id INTEGER NOT NULL,
    exercise_id INTEGER NOT NULL,
    UNIQUE (workout_id, exercise_id),
    FOREIGN KEY (workout_id) REFERENCES workouts(id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
);

-- Create the ExerciseExecution table
CREATE TABLE exercise_executions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    exercise_id INTEGER NOT NULL,
    workout_id INTEGER NOT NULL,
    reps INTEGER NOT NULL,
    execution_date DATETIME NOT NULL,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id),
    FOREIGN KEY (workout_id) REFERENCES workouts(id)
);
