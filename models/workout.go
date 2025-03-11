package models

type Workout struct {
	ID      int       `json:"id"`
	Date    string    `json:"date"`    // (e.g., "01/10/2023")
	TimeIn  string    `json:"time_in"` // (e.g., "10:00")
	TimeOut string    `json:"time_out"` // (e.g., "11:00")
	MoodIn  string    `json:"mood_in"`
	MoodOut string    `json:"mood_out"`
	Lifts   []string  `json:"lifts"`
	Weight  []float64 `json:"weight"`
	Reps    []int     `json:"reps"`
	Sets    []int     `json:"sets"`
}

type Lift struct {
    Name   string
    Weight float64
    Reps   int
    Sets   int
}