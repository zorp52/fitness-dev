package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DbInit() (*sql.DB, error) {
    dbExists := false
    
    if _, err := os.Stat("fitness.db"); !os.IsNotExist(err) {
        dbExists = true
        log.Println("Using existing database file fitness.db")
    } else {
        log.Println("Database file does not exist. Creating fitness.db...")
        file, err := os.Create("fitness.db")
        if err != nil {
            return nil, fmt.Errorf("failed to create database file: %v", err)
        }
        file.Close()
    }

    db, err := sql.Open("sqlite3", "fitness.db")
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }

    if !dbExists {
        dbCreateWorkouts := `
        CREATE TABLE IF NOT EXISTS workouts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            day TEXT NOT NULL,          -- Format: "dd/mm/yyyy"
            time_in TEXT NOT NULL,      -- Format: "HH:MM"
            time_out TEXT NOT NULL,     -- Format: "HH:MM"
            mood_in TEXT NOT NULL,      -- Mood at the start
            mood_out TEXT NOT NULL      -- Mood at the end
        );`
        _, err = db.Exec(dbCreateWorkouts)
        if err != nil {
            return nil, fmt.Errorf("failed to create workouts table: %v", err)
        }

        dbCreateLifts := `
        CREATE TABLE IF NOT EXISTS lifts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            workout_id INTEGER NOT NULL, -- Foreign key to workouts.id
            name TEXT NOT NULL,          -- Exercise name
            weight REAL NOT NULL,        -- Weight in kg
            reps INTEGER NOT NULL,       -- Number of reps
            sets INTEGER NOT NULL,       -- Number of sets
            FOREIGN KEY (workout_id) REFERENCES workouts(id)
        );`
        _, err = db.Exec(dbCreateLifts)
        if err != nil {
            return nil, fmt.Errorf("failed to create lifts table: %v", err)
        }

        log.Println("Database tables created successfully!")
    }

    log.Println("Database initialized successfully!")
    return db, nil
}