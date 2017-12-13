package strength

import "time"
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
  workoutList := WorkoutList {
    Workout { Name: "Squat" },
    Workout { Name: "Deadlift" },
    Workout { Name: "Bench Press" },
    Workout { Name: "Overhead Press" },
    Workout { Name: "Barbell Row" },
  }

  strengthList:= StrengthList {
    Strength {
      Date: time.Now().Unix(),
      WorkoutList: workoutList,
    },
  }

  return strengthList;
}

// var Error = errors.New("Shit fucked yo!")