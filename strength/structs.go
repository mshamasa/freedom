package strength

// Strength is main struct to return for displaying all the rows
type Strength struct {
	Date        int64     `json:"date"`
	WorkoutList []Workout `json:"workoutList"`
}

// Workout is the main struct for a work out
type Workout struct {
	RowID     int32   `gorm:"AUTO_INCREMENT;column:rowid" json:"rowId"`
	UserID    string  `gorm:"column:userId" json:"userId"`
	Exercise  int32   `json:"exercise"`
	Weight    float32 `json:"weight"`
	Sets      int32   `json:"sets"`
	Reps      int32   `json:"reps"`
	Completed int32   `json:"completed"`
	Date      int64   `json:"date"`
}

// Row use to update all rows the new user selected date
type Row struct {
	RowIds []int32 `json:"rowIds"`
	Date   int64   `json:"date"`
}
