package main

import (
	"fmt"
	"log"
	"database/sql"
	"net/http"

	"fitness-dev/api"
	"fitness-dev/backend"
	"fitness-dev/models"
	"fitness-dev/mock"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/inancgumus/screen"
	"github.com/gin-contrib/cors"
)

func main() {
	db, err := backend.DbInit()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	startCLI(db)
}

func startCLI(db *sql.DB) {
	fmt.Scanln() 
	for {
		fmt.Println("Welcome to the Fitness App CLI")
		fmt.Println("1 - Start Server")
		fmt.Println("2 - View/Edit Workouts")
		fmt.Println("3 - Exit")
		fmt.Print("Please enter a number to continue: ")

		var userinput int
		fmt.Scan(&userinput)

		switch userinput {
		case 1:
			startServer(db)
		case 2:
			screen.Clear()
			manageWorkouts(db)
		case 3:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func startServer(db *sql.DB) {
	// Cors & GIN
	router := gin.Default()
	router.Use(cors.Default())

	// Frontend
	router.Static("/frontend", "./fitness-dev/frontend")
	// This means:
	// http://localhost:8080/frontend/index.html will serve fitness-dev/frontend/index.html
	// http://localhost:8080/frontend/style.css will serve fitness-dev/frontend/style.css

	// API routes
	router.POST("/workouts", api.CreateWorkoutHandler(db))         // Create a new workout
	router.GET("/workouts/:day", api.GetWorkoutByDayHandler(db))   // Fetch workout by day
	router.GET("/workouts", api.GetWorkoutsByDateRangeHandler(db)) // Fetch workouts by date range
	router.PUT("/workouts/:id", api.UpdateWorkoutHandler(db))      // Update a workout by ID
	router.DELETE("/workouts/:id", api.DeleteWorkoutHandler(db))   // Delete a workout by ID

	// Default landing page
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Fitness App API!"})
	})
	// favicon aka logo/icon for site
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	

	log.Println("🚀 Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func manageWorkouts(db *sql.DB) {
	fmt.Scanln() 
	for {
		fmt.Println("View/Edit Workouts")
		fmt.Println("1 - View Workouts")
		fmt.Println("2 - Add Workout")
		fmt.Println("3 - Insert Mock Data")
		fmt.Println("4 - Back")
		fmt.Print("Please enter a number to continue: ")

		var userinput int
		fmt.Scan(&userinput)

		switch userinput {
		case 1:
			viewWorkouts(db)
		case 2:
			addWorkout(db)
		case 3:
			mock.InsertMockData(db)
			fmt.Println("Press Enter to continue...")
			fmt.Scanln() 
		case 4:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func viewWorkouts(db *sql.DB) {
	fmt.Print("Enter date (DD-MM-YYYY) to view workout: ")
	var day string
	fmt.Scan(&day)

	workout, err := backend.GetWorkoutByDay(db, day)
	if err != nil {
		fmt.Printf("Failed to fetch workout: %v\n", err)
		return
	}

	fmt.Println("Workout: ", workout)
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() 
}

func addWorkout(db *sql.DB) {
	var workout models.Workout

	fmt.Print("Enter date (DD-MM-YYYY): ")
	fmt.Scan(&workout.Date)

	fmt.Print("Enter time in (HH:MM): ")
	fmt.Scan(&workout.TimeIn)

	fmt.Print("Enter time out (HH:MM): ")
	fmt.Scan(&workout.TimeOut)

	fmt.Print("Enter mood in: ")
	fmt.Scan(&workout.MoodIn)

	fmt.Print("Enter mood out: ")
	fmt.Scan(&workout.MoodOut)

	for {
		var lift models.Lift
		fmt.Print("Enter lift name (or 'done' to finish): ")
		fmt.Scan(&lift.Name)
		if lift.Name == "done" {
			break
		}

		fmt.Print("Enter weight (kg): ")
		fmt.Scan(&lift.Weight)

		fmt.Print("Enter reps: ")
		fmt.Scan(&lift.Reps)

		fmt.Print("Enter sets: ")
		fmt.Scan(&lift.Sets)

		workout.Lifts = append(workout.Lifts, lift.Name)
		workout.Weight = append(workout.Weight, lift.Weight)
		workout.Reps = append(workout.Reps, lift.Reps)
		workout.Sets = append(workout.Sets, lift.Sets)
	}

	if err := backend.InsertWorkout(db, workout); err != nil {
		fmt.Printf("Failed to create workout: %v\n", err)
	} else {
		fmt.Println("Workout created successfully!")
	}

	fmt.Println("Press Enter to continue...")
	fmt.Scanln() 
}