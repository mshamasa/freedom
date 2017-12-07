package strength

import "time"

type Strength struct {
  Due         time.Time `json:"due"`
  WorkoutList []Workout `json:"workoutList"`
}

type Workout struct {
  Name      string    `json:"name"`
  Completed bool      `json:"completed"`
  Sets      int       `json:"sets"`
  Reps      int       `json:"reps"`
}
