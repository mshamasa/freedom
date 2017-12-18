package strength

type Strength struct {
  Date         int64 `json:"date"`
  WorkoutList []Workout `json:"workoutList"`
}

type Workout struct {
  RowId     int32     `gorm:"column:rowid" json:"rowId"`
  UserId    string    `gorm:"column:userId" json:"userId"`
  Exercise  int32     `json:"exercise"`
  Weight    float32   `json:"weight"`
  Sets      int32     `json:"sets"`
  Reps      int32     `json:"reps"`
  Completed int32     `json:"completed"`
  Date      int64     `json:"date"`
}