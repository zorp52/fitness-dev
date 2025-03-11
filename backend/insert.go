package backend

import (
	"database/sql"
	"fmt"
	"fitness-dev/models"
)

func executeInTransaction(db *sql.DB, fn func(tx *sql.Tx) error) error {
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Rollback transaction if error
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Failed to rollback transaction: %v\n", rbErr)
			}
		}
	}()

	// Execute the function
	if err := fn(tx); err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func validateWorkoutForInsert(workout models.Workout) error {
	if workout.Date == "" {
		return fmt.Errorf("date is required")
	}
	if workout.TimeIn == "" {
		return fmt.Errorf("time_in is required")
	}
	if workout.TimeOut == "" {
		return fmt.Errorf("time_out is required")
	}
	if workout.MoodIn == "" {
		return fmt.Errorf("mood_in is required")
	}
	if workout.MoodOut == "" {
		return fmt.Errorf("mood_out is required")
	}
	if len(workout.Lifts) == 0 {
		return fmt.Errorf("at least one lift is required")
	}
	return nil
}

func InsertWorkout(db *sql.DB, workout models.Workout) error {
	if err := validateWorkoutForInsert(workout); err != nil {
		return err
	}

	return executeInTransaction(db, func(tx *sql.Tx) error {
		workoutQuery := `INSERT INTO workouts (day, time_in, time_out, mood_in, mood_out) VALUES (?, ?, ?, ?, ?)`
		result, err := tx.Exec(workoutQuery, workout.Date, workout.TimeIn, workout.TimeOut, workout.MoodIn, workout.MoodOut)
		if err != nil {
			return fmt.Errorf("failed to insert workout: %v", err)
		}

		workoutID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to retrieve workout ID: %v", err)
		}

		liftQuery := `INSERT INTO lifts (workout_id, name, weight, reps, sets) VALUES (?, ?, ?, ?, ?)`
		for i := 0; i < len(workout.Lifts); i++ {
			_, err = tx.Exec(liftQuery, workoutID, workout.Lifts[i], workout.Weight[i], workout.Reps[i], workout.Sets[i])
			if err != nil {
				return fmt.Errorf("failed to insert lift: %v", err)
			}
		}

		return nil
	})
}

func UpdateWorkout(db *sql.DB, workout models.Workout) error {
	return executeInTransaction(db, func(tx *sql.Tx) error {
		workoutQuery := `UPDATE workouts SET `
		var args []interface{}
		if workout.Date != "" {
			workoutQuery += `day = ?, `
			args = append(args, workout.Date)
		}
		if workout.TimeIn != "" {
			workoutQuery += `time_in = ?, `
			args = append(args, workout.TimeIn)
		}
		if workout.TimeOut != "" {
			workoutQuery += `time_out = ?, `
			args = append(args, workout.TimeOut)
		}
		if workout.MoodIn != "" {
			workoutQuery += `mood_in = ?, `
			args = append(args, workout.MoodIn)
		}
		if workout.MoodOut != "" {
			workoutQuery += `mood_out = ?, `
			args = append(args, workout.MoodOut)
		}

		// Remove the trailing comma and space
		workoutQuery = workoutQuery[:len(workoutQuery)-2]
		workoutQuery += ` WHERE id = ?`
		args = append(args, workout.ID)

		_, err := tx.Exec(workoutQuery, args...)
		if err != nil {
			return fmt.Errorf("failed to update workout: %v", err)
		}

		if len(workout.Lifts) > 0 {
			liftDeleteQuery := `DELETE FROM lifts WHERE workout_id = ?`
			_, err = tx.Exec(liftDeleteQuery, workout.ID)
			if err != nil {
				return fmt.Errorf("failed to delete lifts: %v", err)
			}

			liftQuery := `INSERT INTO lifts (workout_id, name, weight, reps, sets) VALUES (?, ?, ?, ?, ?)`
			for i := 0; i < len(workout.Lifts); i++ {
				_, err = tx.Exec(liftQuery, workout.ID, workout.Lifts[i], workout.Weight[i], workout.Reps[i], workout.Sets[i])
				if err != nil {
					return fmt.Errorf("failed to insert lift: %v", err)
				}
			}
		}

		return nil
	})
}

func DeleteWorkout(db *sql.DB, workoutID int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	workoutDeleteQuery := `DELETE FROM workouts WHERE id = ?`
	_, err = tx.Exec(workoutDeleteQuery, workoutID)
	if err != nil {
		return fmt.Errorf("failed to delete workout: %v", err)
	}

	liftDeleteQuery := `DELETE FROM lifts WHERE workout_id = ?`
	_, err = tx.Exec(liftDeleteQuery, workoutID)
	if err != nil {
		return fmt.Errorf("failed to delete lifts: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

