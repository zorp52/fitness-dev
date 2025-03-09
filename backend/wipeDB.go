package database

import (
	"database/sql"
	"fmt"
)

// WipeDB deletes all data from the database
func WipeDB(db *sql.DB) error {
	queries := []string{
		`DELETE FROM workouts`,
		`DELETE FROM lifts`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("failed to wipe database: %v", err)
		}
	}

	return nil
}
