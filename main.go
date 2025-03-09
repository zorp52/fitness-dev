package main

import (
	"database/sql"
	"log"
	"fitness-dev/api"
	"fitness-dev/database"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.DbInit()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Cors & GIN
	router := gin.Default()
	router.Use(cors.Default())

	// API routes
	router.POST("/workouts", api.CreateWorkoutHandler(db))         // Create a new workout
	router.GET("/workouts/:day", api.GetWorkoutByDayHandler(db))   // Fetch workout by day
	router.GET("/workouts", api.GetWorkoutsByDateRangeHandler(db)) // Fetch workouts by date range
	router.PUT("/workouts/:id", api.UpdateWorkoutHandler(db))      // Update a workout by ID
	router.DELETE("/workouts/:id", api.DeleteWorkoutHandler(db))   // Delete a workout by ID

	// Start server
	log.Println("🚀 Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}