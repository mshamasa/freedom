package strength

type Strength struct {
  Date         int64 `json:"date"`
  WorkoutList []Workout `json:"workoutList"`
}

type Workout struct {
  Name      string    `json:"name"`
  Completed bool      `json:"completed"`
  Sets      int       `json:"sets"`
  Reps      int       `json:"reps"`
}
