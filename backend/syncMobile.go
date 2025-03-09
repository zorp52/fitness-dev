package database

import (
    "database/sql"
    "fmt"
    "encoding/json"
    "fitness-dev/models"
)

func SyncMobile(db *sql.DB, jsonData []byte) error {
    var workout models.Workout
    err := json.Unmarshal(jsonData, &workout)
    if err != nil {
        return fmt.Errorf("failed to extract data from json: %v", err)
    }

    return InsertWorkout(db, workout)
}