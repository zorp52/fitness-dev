package models

import "time"

type Workout struct {
    ID       int
    Date     time.Time
    TimeIn   time.Time
    TimeOut  time.Time
    MoodIn   string
    MoodOut  string
    Lifts    []string
    Weight   []float64
    Reps     []int
    Sets     []int
}

type Lift struct {
    Name   string
    Weight float64
    Reps   int
    Sets   int
}