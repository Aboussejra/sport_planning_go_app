-- Create Exercises table
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create ExerciseExecutions table
CREATE TABLE exercise_executions (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL,
    reps INT NOT NULL,
    execution_date DATE NOT NULL,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
);

-- Create Workouts table
CREATE TABLE workouts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    day_of_week VARCHAR(20) NOT NULL
);