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

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/workouts/:day", api.GetWorkoutByDayHandler(db)) // Fetch workout by day
	router.GET("/workouts", api.GetWorkoutsByDateRangeHandler(db)) // Fetch workouts by date range

	log.Println("🚀 Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}