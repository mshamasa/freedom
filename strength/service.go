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
  Add()
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

func (strengthService) Add() {

}


// var Error = errors.New("Shit fucked yo!")