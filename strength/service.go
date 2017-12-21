package strength

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// what the flying fuck with these comments.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// import "errors"

// StrengthService interface that will require implementation of methods.
type StrengthService interface {
	//TODO error handling...later
	// Index() (StrengthList, error)
	Index(request interface{}) StrengthList
	AddRow(request interface{})
	SaveRow(request interface{})
	SaveWorkout(request interface{}) Workout
	UpdateRowsDate(request interface{})
	DeleteRow(request interface{})
}

type strengthService struct{}

type StrengthList []Strength
type WorkoutList []Workout

func (strengthService) Index(request interface{}) StrengthList {
	workouts := WorkoutList{}
	db, err := gorm.Open("sqlite3", "./strength.db")
	if err != nil {
		fmt.Println("could not connect to database.")
	}
	defer db.Close()
	// query
	req := request.(strengthRequest)
	// TODO set limit based on dates instead of this way
	db.Raw("SELECT rowid, * FROM workouts WHERE userId = ?", req.UserID).Order("date desc").Limit(20).Scan(&workouts)
	db.Close()

	strengthList := sortWorkouts(workouts)

	return strengthList
}

func (strengthService) AddRow(request interface{}) {
	req := request.(strengthRequest)

	db, err := gorm.Open("sqlite3", "./strength.db")
	if err != nil {
		fmt.Println("could not connect to database.")
	}
	defer db.Close()
	// create 3 default rows
	for i := 0; i < 3; i++ {
		workout := Workout{
			UserID:   req.UserID,
			Exercise: int32(i + 1),
			Date:     req.StartDate,
		}
		db.Create(&workout)
	}
	db.Close()
}

func (strengthService) SaveRow(request interface{}) {
	req := request.(strengthRequest)
	UserID := req.UserID
	strengthList := req.List

	db, err := gorm.Open("sqlite3", "./strength.db")
	if err != nil {
		fmt.Println("could not connect to database.")
	}
	defer db.Close()

	for i := 0; i < len(strengthList); i++ {
		date := strengthList[i].Date
		workouts := generateWorkouts(date, UserID, strengthList[i].WorkoutList)
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

	db.Model(&workout).Where("rowid = ?", workout.RowID).Updates(map[string]interface{}{
		"exercise":  workout.Exercise,
		"weight":    workout.Weight,
		"sets":      workout.Sets,
		"reps":      workout.Reps,
		"completed": workout.Completed,
		"date":      workout.Date,
	})
	db.Raw("SELECT rowid, * from workouts WHERE rowid =?", workout.RowID).Scan(&wk)
	db.Close()
	return wk
}

func (strengthService) UpdateRowsDate(request interface{}) {
	req := request.(strengthRequest)
	row := req.Row

	db, err := gorm.Open("sqlite3", "./strength.db")
	if err != nil {
		fmt.Println("could not connect to database.")
	}
	defer db.Close()

	db.Table("workouts").Where("rowid IN (?)", row.RowIds).Updates(map[string]interface{}{
		"date": row.Date,
	})
	db.Close()
}

func (strengthService) DeleteRow(request interface{}) {
	req := request.(strengthRequest)
	row := req.Row

	db, err := gorm.Open("sqlite3", "./strength.db")
	if err != nil {
		fmt.Println("could not connect to database.")
	}
	defer db.Close()

	db.Debug().Table("workouts").Where("rowid IN (?)", row.RowIds).Delete(&Workout{})
	db.Close()
}

// var Error = errors.New("Shit fucked yo!")
