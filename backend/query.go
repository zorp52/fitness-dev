package database

import (
	"database/sql"
	"fmt"
	"fitness-dev/models"
)

// GetWorkoutByDay fetches a workout for a specific day.
func GetWorkoutByDay(db *sql.DB, day string) (models.Workout, error) {
	query := `SELECT id, day, time_in, time_out, mood_in, mood_out FROM workouts WHERE day = ?`
	row := db.QueryRow(query, day)

	var workout models.Workout
	err := row.Scan(&workout.ID, &workout.Date, &workout.TimeIn, &workout.TimeOut, &workout.MoodIn, &workout.MoodOut)
	if err != nil {
		return models.Workout{}, fmt.Errorf("failed to fetch workout: %v", err)
	}

	// Fetch lifts for this workout
	query = `SELECT name, weight, reps, sets FROM lifts WHERE workout_id = ?`
	rows, err := db.Query(query, workout.ID)
	if err != nil {
		return models.Workout{}, fmt.Errorf("failed to fetch lifts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var lift models.Lift
		if err := rows.Scan(&lift.Name, &lift.Weight, &lift.Reps, &lift.Sets); err != nil {
			return models.Workout{}, fmt.Errorf("failed to scan row: %v", err)
		}
		workout.Lifts = append(workout.Lifts, lift.Name)
		workout.Weight = append(workout.Weight, lift.Weight)
		workout.Reps = append(workout.Reps, lift.Reps)
		workout.Sets = append(workout.Sets, lift.Sets)
	}

	return workout, nil
}

// GetWorkoutsByDateRange fetches workouts within a date range.
func GetWorkoutsByDateRange(db *sql.DB, startDate, endDate string) ([]models.Workout, error) {
	query := `SELECT id, day, time_in, time_out, mood_in, mood_out FROM workouts WHERE day BETWEEN ? AND ?`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch workouts: %v", err)
	}
	defer rows.Close()

	var workouts []models.Workout
	for rows.Next() {
		var workout models.Workout
		if err := rows.Scan(&workout.ID, &workout.Date, &workout.TimeIn, &workout.TimeOut, &workout.MoodIn, &workout.MoodOut); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Fetch lifts for this workout
		liftsQuery := `SELECT name, weight, reps, sets FROM lifts WHERE workout_id = ?`
		liftsRows, err := db.Query(liftsQuery, workout.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch lifts: %v", err)
		}
		defer liftsRows.Close()

		for liftsRows.Next() {
			var lift models.Lift
			if err := liftsRows.Scan(&lift.Name, &lift.Weight, &lift.Reps, &lift.Sets); err != nil {
				return nil, fmt.Errorf("failed to scan lift row: %v", err)
			}
			workout.Lifts = append(workout.Lifts, lift.Name)
			workout.Weight = append(workout.Weight, lift.Weight)
			workout.Reps = append(workout.Reps, lift.Reps)
			workout.Sets = append(workout.Sets, lift.Sets)
		}

		workouts = append(workouts, workout)
	}

	return workouts, nil
}