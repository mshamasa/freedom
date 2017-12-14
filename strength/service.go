package strength

import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
)
// import "errors"

// StrengthService handles returning info
type StrengthService interface {
  //TODO error handling...later
  // Index() (StrengthList, error)
  Index() StrengthList
}

type strengthService struct{}

type StrengthList []Strength
type WorkoutList []Workout

func (strengthService) Index() StrengthList {
  workoutList := WorkoutList {}

  db, err := sql.Open("sqlite3", "./strength.db")
  checkErr(err)
  if err != nil {
    fmt.Println("could not connect to database.")
  }
  defer db.Close()
  // query
  rows, err := db.Query("SELECT * FROM workouts")
  checkErr(err)
  var userId string
  var exercise int32
  var weight float32
  var sets int32
  var reps int32
  var completed int32
  var date int32

  for rows.Next() {
    err = rows.Scan(&userId, &exercise, &weight, &sets, &reps, &completed, &date)
    checkErr(err)

    workout := Workout {
      UserId: userId,
      Exercise: exercise,
      Weight: weight,
      Sets: sets,
      Reps: reps,
      Completed: completed,
      Date: date,
    }
    workoutList = append(workoutList, workout)
  }

  rows.Close()
  db.Close()

  strengthList:= StrengthList {
    Strength {
      WorkoutList: workoutList,
    },
  }

  return strengthList;
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

// var Error = errors.New("Shit fucked yo!")