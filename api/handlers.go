package api

import (
	"net/http"
	"fitness-dev/database"
	"fitness-dev/models"

	"github.com/gin-gonic/gin"
)

func GetWorkoutByDayHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		day := c.Param("day") // Get the day from the URL parameter

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
		startDate := c.Query("startDate") // Get start date from query parameters
		endDate := c.Query("endDate")     // Get end date from query parameters

		// validate date range
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