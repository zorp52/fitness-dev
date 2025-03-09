package api

import (
	"net/http"
	"fitness-dev/database"
	"fitness-dev/models"

	"github.com/gin-gonic/gin"
)

func CreateWorkoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var workout models.Workout
		if err := c.ShouldBindJSON(&workout); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.InsertWorkout(db, workout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout created successfully"})
	}
}

func GetWorkoutByDayHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		day := c.Param("day")
		workout, err := database.GetWorkoutByDay(db, day)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, workout)
	}
}

func GetWorkoutsByDateRangeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate := c.Query("startDate")
		endDate := c.Query("endDate")

		if startDate == "" || endDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
			return
		}

		workouts, err := database.GetWorkoutsByDateRange(db, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, workouts)
	}
}

func UpdateWorkoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var workout models.Workout
		if err := c.ShouldBindJSON(&workout); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		workout.ID = id // Set the workout ID from the URL parameter
		if err := database.UpdateWorkout(db, workout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout updated successfully"})
	}
}

func DeleteWorkoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := database.DeleteWorkout(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout deleted successfully"})
	}
}