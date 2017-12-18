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
  Save(request interface{})
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
  db.Where("userId = ?", req.UserId).Order("date desc").Limit(20).Find(&workouts)
  db.Close()

  strengthList := sortWorkouts(workouts)

  return strengthList;
}

func (strengthService) Save(request interface{}) {
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


// var Error = errors.New("Shit fucked yo!")