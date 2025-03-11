package backend

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DbInit() (*sql.DB, error) {
	dbExists := false
	
	// Check if the file exists
	if _, err := os.Stat("fitness.db"); !os.IsNotExist(err) {
		dbExists = true
		log.Println("Using existing database file fitness.db")
	} else {
		log.Println("Database file does not exist. Creating fitness.db...")
	}

	// Open or create the database
	db, err := sql.Open("sqlite3", "fitness.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Ensure the database file is actually created
	if !dbExists {
		file, err := os.Create("fitness.db")
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close()
	}

	// Enable foreign keys (important for referencing between tables)
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Printf("Failed to enable foreign keys: %v", err)
	}

	// Check if the table exists
	var tableCount int
	err = db.QueryRow("SELECT count(name) FROM sqlite_master WHERE type='table' AND name IN ('workouts', 'lifts')").Scan(&tableCount)
	if err != nil {
		log.Printf("Error checking tables: %v", err)
	}

	// If tables don't exist, create them
	if tableCount == 0 {
		log.Println("Creating tables...")

		dbCreateWorkouts := `
		CREATE TABLE IF NOT EXISTS workouts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			day TEXT NOT NULL,         
			time_in TEXT NOT NULL,     
			time_out TEXT NOT NULL,    
			mood_in TEXT NOT NULL,     
			mood_out TEXT NOT NULL     
		);`
		_, err = db.Exec(dbCreateWorkouts)
		if err != nil {
			return nil, fmt.Errorf("failed to create workouts table: %v", err)
		}

		dbCreateLifts := `
		CREATE TABLE IF NOT EXISTS lifts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			workout_id INTEGER NOT NULL, 
			name TEXT NOT NULL,          
			weight REAL NOT NULL,        
			reps INTEGER NOT NULL,       
			sets INTEGER NOT NULL,       
			FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE
		);`
		_, err = db.Exec(dbCreateLifts)
		if err != nil {
			return nil, fmt.Errorf("failed to create lifts table: %v", err)
		}

		log.Println("Database tables created successfully!")
	} else {
		log.Println("Tables already exist. Skipping creation.")
	}

	log.Println("Database initialized successfully!")
	return db, nil
}
