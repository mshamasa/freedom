package strength

import (
  "fmt"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)
// import "errors"

// StrengthService handles returning info
type StrengthService interface {
  //TODO error handling...later
  // Index() (StrengthList, error)
  Index(request interface{}) StrengthList
  SaveRow(request interface{})
  SaveWorkout(request interface{}) Workout
}

type strengthService struct{}

type StrengthList []Strength
type WorkoutList []Workout

func (strengthService) Index(request interface{}) StrengthList {
  workouts := WorkoutList {}
  db, err := gorm.Open("sqlite3", "./strength.db")
  if err != nil {
    fmt.Println("could not connect to database.")
  }
  defer db.Close()
  // query
  req := request.(strengthRequest)
  // TODO set limit based on dates instead of this way
  db.Raw("SELECT rowid, * FROM workouts WHERE userId = ?", req.UserId).Order("date desc").Limit(20).Scan(&workouts)
  db.Close()

  strengthList := sortWorkouts(workouts)

  return strengthList;
}

func (strengthService) SaveRow(request interface{}) {
  req := request.(strengthRequest)
  userId := req.UserId
  strengthList := req.List

  db, err := gorm.Open("sqlite3", "./strength.db")
  if err != nil {
    fmt.Println("could not connect to database.")
  }
  defer db.Close()

  for i := 0; i < len(strengthList); i++ {
    date := strengthList[i].Date
    workouts := generateWorkouts(date, userId, strengthList[i].WorkoutList)
    for j := 0; j < len(workouts); j++ {
      wk := workouts[j]
      db.Create(&wk)
    }
  }
  db.Close()
}

func (strengthService) SaveWorkout(request interface{}) Workout {
  var wk Workout

  req := request.(strengthRequest)
  workout := req.Workout

  db, err := gorm.Open("sqlite3", "./strength.db")
  if err != nil {
    fmt.Println("could not connect to database.")
  }
  defer db.Close()

  db.Model(&workout).Where("rowid = ?", workout.RowId).Updates(map[string]interface{}{
    "exercise": workout.Exercise,
    "weight": workout.Weight,
    "sets": workout.Sets,
    "reps": workout.Reps,
    "completed": workout.Completed,
    "date": workout.Date,
  })
  db.Raw("SELECT rowid, * from workouts WHERE rowid =?", workout.RowId).Scan(&wk)
  db.Close()
  return wk
}


// var Error = errors.New("Shit fucked yo!")