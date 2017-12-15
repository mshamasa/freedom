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
  db.Where("userId = ?", req.UserId).Order("date desc").Find(&workouts)
  db.Close()

  workoutsMap := make(map[int32][]Workout)
  for i := 0; i < len(workouts); i++ {
    date := workouts[i].Date
    workoutsMap[date] = append(workoutsMap[date], workouts[i])
  }

  strengthList := StrengthList{}
  for k, v := range workoutsMap {
    strength := Strength {
      Date: k,
      WorkoutList: v,
    }
    strengthList = append(strengthList, strength)
  }

  return strengthList;
}


// var Error = errors.New("Shit fucked yo!")