package strength

import (
	"database/sql"
	"fmt"
	"log"
	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

// List slice of Strength objects
type List []Strength

// WorkoutList slice of Workout objects
type WorkoutList []Workout

// Index will get all user data with limit set to 20 rows
func Index(req Request) List {
	workoutList := WorkoutList{}
	db := getDatabaseConnection()

	rows, err := db.Query("SELECT * FROM workouts WHERE userID=? ORDER BY date DESC LIMIT 20", req.UserID)
	if err != nil {
		log.Fatalf("ERROR IN QUERY!: %v", err)
	}

	for rows.Next() {
		var row Workout
		if err := rows.Scan(&row.RowID, &row.UserID, &row.Exercise, &row.Weight,
			&row.Sets, &row.Reps, &row.Completed, &row.Date); err != nil {
			log.Fatalf("ERROR IN ROWS!: %v", err)
		}
		workoutList = append(workoutList, row)
	}
	defer rows.Close()
	defer db.Close()

	strengthList := sortWorkouts(workoutList)

	return strengthList
}

// DeleteRow will delete rows from db
func DeleteRow(req Request) {
	row := req.Row

	db := getDatabaseConnection()
	stmt, err := db.Prepare("DELETE FROM workouts WHERE rowID=?")
	checkErr(err)
	rowIds := row.RowIds
	for i := 0; i < len(rowIds); i++ {
		stmt.Exec(rowIds[i])
	}
	db.Close()
}

// AddRows will add row to db
func AddRows(req Request) {
	db := getDatabaseConnection()
	stmt, err := db.Prepare("INSERT INTO workouts (userID, exercise, date) VALUES (?,?,?)")
	checkErr(err)

	// create 3 default rows
	for i := 0; i < req.Amount; i++ {
		res, err := stmt.Exec(req.UserID, (i + 1), req.StartDate)
		checkErr(err)
		affect, err := res.RowsAffected()
		checkErr(err)
		log.Println(affect)
	}
	defer db.Close()
}

// UpdateRowsDate update each row with new date
func UpdateRowsDate(req Request) {
	row := req.Row

	db := getDatabaseConnection()
	stmt, err := db.Prepare("UPDATE workouts SET date=? WHERE rowID=?")
	checkErr(err)
	rowIds := row.RowIds
	for i := 0; i < len(rowIds); i++ {
		stmt.Exec(row.Date, rowIds[i])
	}
	db.Close()
}

// SaveWorkout will update a workout in db
func SaveWorkout(req Request) Workout {
	workout := req.Workout
	db := getDatabaseConnection()

	stmt, err := db.Prepare("UPDATE workouts SET exercise=?, weight=?, sets=?, reps=?, completed=?, date=? WHERE rowID=?")
	checkErr(err)
	res, err := stmt.Exec(
		workout.Exercise,
		workout.Weight,
		workout.Sets,
		workout.Reps,
		workout.Completed,
		workout.Date,
		workout.RowID,
	)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	log.Println(affect)

	rows, err := db.Query("SELECT * FROM workouts where rowID=?", workout.RowID)
	checkErr(err)

	var wk Workout
	for rows.Next() {
		if err := rows.Scan(&wk.RowID, &wk.UserID, &wk.Exercise, &wk.Weight,
			&wk.Sets, &wk.Reps, &wk.Completed, &wk.Date); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
	db.Close()

	return wk
}

func checkErr(err error) {
	if err != nil {
		log.Panicf("%s Error: ", err)
	}
}

func getDatabaseConnection() *sql.DB {
	var db *sql.DB
	var err error
	var connectionString string

	connectionString = fmt.Sprintf("rootuser:!6rootuser@tcp(strengthdbinstance.c6edl7uupl3x.us-east-1.rds.amazonaws.com)/strength")

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	return db
}
