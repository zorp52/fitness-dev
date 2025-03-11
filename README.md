# Fitness App Backend Documentation

This documentation provides an overview of the backend functionality for the Fitness App. It includes details on how to interact with the API, the data models, and the database structure. This guide is intended for frontend developers who need to integrate with the backend.

---

## Table of Contents

1. **API Endpoints**
   - Create Workout
   - Get Workout by Day
   - Get Workouts by Date Range
   - Update Workout
   - Delete Workout

2. **Data Models**
   - Workout
   - Lift

3. **Database Structure**
   - Workouts Table
   - Lifts Table

4. **Mock Data**
   - Inserting Mock Data

5. **Error Handling**
   - Common Errors

6. **Folder Structure**
   - Overview of the project structure

---

## 1. API Endpoints

The backend provides the following RESTful API endpoints for managing workouts:

### 1.1 Create Workout
- **Endpoint**: `POST /workouts`
- **Description**: Creates a new workout.
- **Request Body**:
  ```json
  {
    "date": "01/10/2023",
    "time_in": "10:00",
    "time_out": "11:00",
    "mood_in": "Good",
    "mood_out": "Great",
    "lifts": ["Squat", "Bench"],
    "weight": [100.0, 80.0],
    "reps": [5, 10],
    "sets": [3, 4]
  }
  ```
- **Response**:
  ```json
  {
    "message": "Workout created successfully"
  }
  ```

### 1.2 Get Workout by Day
- **Endpoint**: `GET /workouts/:day`
- **Description**: Retrieves a workout by a specific day.
- **URL Parameter**: `day` (e.g., `01/10/2023`)
- **Response**:
  ```json
  {
    "id": 1,
    "date": "01/10/2023",
    "time_in": "10:00",
    "time_out": "11:00",
    "mood_in": "Good",
    "mood_out": "Great",
    "lifts": ["Squat", "Bench"],
    "weight": [100.0, 80.0],
    "reps": [5, 10],
    "sets": [3, 4]
  }
  ```

### 1.3 Get Workouts by Date Range
- **Endpoint**: `GET /workouts`
- **Description**: Retrieves workouts within a specified date range.
- **Query Parameters**:
  - `startDate`: Start date of the range (e.g., `01/10/2023`)
  - `endDate`: End date of the range (e.g., `01/15/2023`)
- **Response**:
  ```json
  [
    {
      "id": 1,
      "date": "01/10/2023",
      "time_in": "10:00",
      "time_out": "11:00",
      "mood_in": "Good",
      "mood_out": "Great",
      "lifts": ["Squat", "Bench"],
      "weight": [100.0, 80.0],
      "reps": [5, 10],
      "sets": [3, 4]
    },
    {
      "id": 2,
      "date": "01/11/2023",
      "time_in": "09:00",
      "time_out": "10:00",
      "mood_in": "Tired",
      "mood_out": "Energetic",
      "lifts": ["Deadlift", "Press"],
      "weight": [120.0, 60.0],
      "reps": [3, 8],
      "sets": [5, 3]
    }
  ]
  ```

### 1.4 Update Workout
- **Endpoint**: `PUT /workouts/:id`
- **Description**: Updates an existing workout by ID.
- **URL Parameter**: `id` (e.g., `1`)
- **Request Body**:
  ```json
  {
    "date": "01/10/2023",
    "time_in": "10:00",
    "time_out": "11:00",
    "mood_in": "Good",
    "mood_out": "Great",
    "lifts": ["Squat", "Bench"],
    "weight": [100.0, 80.0],
    "reps": [5, 10],
    "sets": [3, 4]
  }
  ```
- **Response**:
  ```json
  {
    "message": "Workout updated successfully"
  }
  ```

### 1.5 Delete Workout
- **Endpoint**: `DELETE /workouts/:id`
- **Description**: Deletes a workout by ID.
- **URL Parameter**: `id` (e.g., `1`)
- **Response**:
  ```json
  {
    "message": "Workout deleted successfully"
  }
  ```

---

## 2. Data Models

### 2.1 Workout
The `Workout` model represents a workout session. It includes the following fields:

```go
type Workout struct {
    ID      int       `json:"id"`
    Date    string    `json:"date"`    // Format: "DD/MM/YYYY"
    TimeIn  string    `json:"time_in"` // Format: "HH:MM"
    TimeOut string    `json:"time_out"` // Format: "HH:MM"
    MoodIn  string    `json:"mood_in"`
    MoodOut string    `json:"mood_out"`
    Lifts   []string  `json:"lifts"`  // List of lift names
    Weight  []float64 `json:"weight"` // List of weights (kg)
    Reps    []int     `json:"reps"`   // List of repetitions
    Sets    []int     `json:"sets"`   // List of sets
}
```

### 2.2 Lift
The `Lift` model represents a single exercise within a workout:

```go
type Lift struct {
    Name   string
    Weight float64
    Reps   int
    Sets   int
}
```

---

## 3. Database Structure

### 3.1 Workouts Table
The `workouts` table stores workout sessions:

| Column    | Type    | Description                     |
|-----------|---------|---------------------------------|
| id        | INTEGER | Primary key, auto-incrementing  |
| day       | TEXT    | Date of the workout (DD/MM/YYYY)|
| time_in   | TEXT    | Start time of the workout (HH:MM)|
| time_out  | TEXT    | End time of the workout (HH:MM) |
| mood_in   | TEXT    | Mood at the start of the workout|
| mood_out  | TEXT    | Mood at the end of the workout  |

### 3.2 Lifts Table
The `lifts` table stores individual exercises within a workout:

| Column     | Type    | Description                     |
|------------|---------|---------------------------------|
| id         | INTEGER | Primary key, auto-incrementing  |
| workout_id | INTEGER | Foreign key referencing workouts|
| name       | TEXT    | Name of the lift                |
| weight     | REAL    | Weight lifted (kg)              |
| reps       | INTEGER | Number of repetitions           |
| sets       | INTEGER | Number of sets                  |

---

## 4. Mock Data

You can insert mock data into the database for testing purposes. This will generate 10 random workouts with random lifts.

- **CLI Command**: Select option `3 - Insert Mock Data` in the CLI.
- **API Endpoint**: Not available via API, only through CLI.

---

## 5. Error Handling

### 5.1 Common Errors
- **400 Bad Request**: Invalid input data (e.g., missing fields, invalid date format).
- **404 Not Found**: Workout not found for the given day or ID.
- **500 Internal Server Error**: Database or server-side error.

---

## 6. Folder Structure

The project is organized as follows:

```
fitness-dev/
├── fitness.db            # SQLite database file
├── go.mod                # Go module file
├── go.sum                # Go dependencies checksum file
├── main.go               # Main application entry point
├── api/
│   └── handlers.go       # API request handlers
├── backend/
│   ├── initDB.go         # Database initialization
│   ├── insert.go         # Workout insertion logic
│   ├── query.go          # Workout query logic
│   ├── syncMobile.go     # Mobile sync functionality
│   └── wipeDB.go         # Database wipe functionality
├── mock/
│   └── mockData.go       # Mock data generation
└── models/
    └── workout.go        # Data models (Workout and Lift)
```
