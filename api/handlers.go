package api

import (
	"net/http"
	"database/sql"
	"strconv"

	"fitness-dev/backend"
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

		if err := backend.InsertWorkout(db, workout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout created successfully"})
	}
}

func GetWorkoutByDayHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		day := c.Param("day")
		workout, err := backend.GetWorkoutByDay(db, day)
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

		workouts, err := backend.GetWorkoutsByDateRange(db, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, workouts)
	}
}

func UpdateWorkoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get workoutID(str) -> workoutID(int)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
			return
		}

		var workout models.Workout
		if err := c.ShouldBindJSON(&workout); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		workout.ID = id
		if err := backend.UpdateWorkout(db, workout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout updated successfully"})
	}
}

func DeleteWorkoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
			return
		}

		if err := backend.DeleteWorkout(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Workout deleted successfully"})
	}
}