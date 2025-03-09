package mock

import (
	"fmt"
	"math/rand"
	"time"
	"fitness-dev/database"
	"fitness-dev/models"

)

fucn devDisplayGeneratedData( workouts []models.Workout) {
	for _, workout := range workouts {
		fmt.Printf("Date: %s\n", workout.Date)
		fmt.Printf("Time In: %s\n", workout.TimeIn)
		fmt.Printf("Time Out: %s\n", workout.TimeOut)
		fmt.Printf("Mood In: %s\n", workout.MoodIn)
		fmt.Printf("Mood Out: %s\n", workout.MoodOut)
		fmt.Println("Lifts:")
		for i := 0; i < len(workout.Lifts); i++ {
			fmt.Printf("  %s: %.2fkg x %d x %d\n", workout.Lifts[i], workout.Weight[i], workout.Reps[i], workout.Sets[i])
		}
		fmt.Println()
	}
}

var moods = [...]string{"Exhausted", "Tired", "Meh", "Good", "Great", "Energetic"}
var lifts = [...]string{"Squat", "Bench", "Deadlift", "Press", "Curls", "Lat Pulldown", "Leg Press", "Leg Curl", "Leg Extension", "Tricep Extension", "Tricep Pushdown", "Tricep Dip", "Bicep Curl", "Bicep Hammer Curl", "Bicep Concentration Curl", "Bicep Preacher Curl", "Bicep Reverse Curl", "Bicep Cable Curl", "Bicep Barbell Curl", "Bicep Dumbbell Curl", "Bicep EZ Curl", "Bicep Incline"}

func insertMockData(db *sql.DB) {
    fmt.Println("Generating mock data...")
    for i := 0; i < 10; i++ {
        workout := models.Workout{
            Date:    time.Now().AddDate(0, 0, i).Format("02/01/2006")
            TimeIn:  time.Now().Format("15:04"),
            TimeOut: time.Now().Add(time.Hour).Format("15:04"),
            MoodIn:  moods[rand.Intn(len(moods))],
            MoodOut: moods[rand.Intn(len(moods))],
            Lifts:  make([]string, 5),
            Weight: make([]float64, 5),
            Reps:   make([]int, 5),
            Sets:   make([]int, 5),
        }
        for j := 0; j < 5; j++ {
            workout.Lifts[j] = lifts[rand.Intn(len(lifts))]
            workout.Weight[j] = float64(rand.Intn(100))
            workout.Reps[j] = rand.Intn(10) + 1
            workout.Sets[j] = rand.Intn(5) + 1
        }

        if err := database.InsertWorkout(db, workout); err != nil {
            fmt.Printf("Failed to insert workout: %v\n", err)
        }
    }
    fmt.Println("Mock data generated.")
	devDisplayGeneratedData(workouts)
}
