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
    optional boolean,
    can_be_replaced boolean,
    UNIQUE (workout_id, exercise_id),
    FOREIGN KEY (workout_id) REFERENCES workouts(id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
);

CREATE TABLE exercise_replacement (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    exercise_id INTEGER NOT NULL,
    replacement_exercise_id INTEGER NOT NULL,
    UNIQUE (exercise_id, replacement_exercise_id)
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

-- Insert example data into exercises table
INSERT INTO exercises (id, name) VALUES (1, 'Développé couché');
INSERT INTO exercises (id, name) VALUES (2, 'Ecarté couché');
INSERT INTO exercises (id, name) VALUES (3, 'Développé épaules haltères');
INSERT INTO exercises (id, name) VALUES (4, 'Traction pronation');
INSERT INTO exercises (id, name) VALUES (5, 'Curl marteau');
INSERT INTO exercises (id, name) VALUES (6, 'Tennis de table');
INSERT INTO exercises (id, name) VALUES (7, 'Traction Supination');

-- Insert example data into workouts table
INSERT INTO workouts (name, day_of_week) VALUES ('Chest Day', 'Monday');
INSERT INTO workouts (name, day_of_week) VALUES ('Leg Day', 'Tuesday');
INSERT INTO workouts (name, day_of_week) VALUES ('Back Day', 'Wednesday');
INSERT INTO workouts (name, day_of_week) VALUES ('Shoulder Day', 'Thursday');
INSERT INTO workouts (name, day_of_week) VALUES ('Arm Day', 'Friday');


-- Insert example data into workout_exercises table
INSERT INTO workout_exercises (workout_id, exercise_id, optional, can_be_replaced) 
VALUES 
(1, 1, 0, 0),  -- Chest Day includes Développé couché (not optional, cannot be replaced)
(1, 2, 0, 0),  -- Chest Day includes Ecarté couché (not optional, cannot be replaced)
(2, 6, 0, 0),  -- Leg Day includes Tennis de table (not optional, cannot be replaced)
(3, 4, 0, 0),  -- Back Day includes Traction pronation (not optional, cannot be replaced)
(4, 3, 0, 1),  -- Shoulder Day includes Développé épaules haltères (not optional, cannot be replaced)
(5, 5, 1, 1);  -- Arm Day includes Curl marteau (optional, can be replaced)

-- Insert example data into exercise_replacement table
INSERT INTO exercise_replacement (exercise_id, replacement_exercise_id) 
VALUES 
(5, 7); -- Curl marteau can be replaced with Traction Supination