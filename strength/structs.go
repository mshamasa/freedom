package strength

// Strength is main struct to return for displaying all the rows
type Strength struct {
	Date        int64     `json:"date"`
	WorkoutList []Workout `json:"workoutList"`
}

// Workout is the main struct for a work out
type Workout struct {
	RowID     int32   `json:"rowID"`
	UserID    string  `json:"userID"`
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

// Request used for all requests
type Request struct {
	UserID    string     `json:"userID"`
	Workout   Workout    `json:"workout"`
	List      []Strength `json:"list"`
	Row       Row        `json:"row"`
	StartDate int64      `json:"startDate"`
	EndDate   int64      `json:"endDate"`
	Amount    int        `json:"amount"`
}

// Response struct for all responses
type Response struct {
	List    []Strength `json:"list"`
	Workout Workout    `json:"workout"`
	Err     string     `json:"err, omitempty"`
	Code    string     `json:"code, omitempty"`
}
