package strength

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
	// register some standard stuff
	_ "google.golang.org/appengine/cloudsql"
)

// Service interface that will have CRUD operations for strenght database
type Service interface {
	//TODO error handling...later
	Index(request interface{}) List
	AddRows(request interface{})
	SaveWorkout(request interface{}) Workout
	UpdateRowsDate(request interface{})
	DeleteRow(request interface{})
}

type strengthService struct{}

// List slice of Strength objects
type List []Strength

// WorkoutList slice of Workout objects
type WorkoutList []Workout

func (strengthService) Index(request interface{}) List {
	req := request.(strengthRequest)
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

func (strengthService) AddRows(request interface{}) {
	req := request.(strengthRequest)
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

func (strengthService) SaveWorkout(request interface{}) Workout {
	req := request.(strengthRequest)
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

func (strengthService) UpdateRowsDate(request interface{}) {
	req := request.(strengthRequest)
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

func (strengthService) DeleteRow(request interface{}) {
	req := request.(strengthRequest)
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

func checkErr(err error) {
	if err != nil {
		log.Panicf("%s Error: ", err)
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}

func getDatabaseConnection() *sql.DB {
	var (
		/* To connect in dev:
		** 1: run './cloud_sql_proxy -instances=freedom-190400:us-central1:strengthworkouts=tcp:3306'
		** 2: then run app in dev
		 */
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user           = mustGetenv("CLOUDSQL_USER")
		password       = mustGetenv("CLOUDSQL_PASSWORD")
		dbName         = mustGetenv("CLOUDSQL_DATABASE")
		devConnection  = mustGetenv("SQL_DEV_CONNECTION")
	)

	var db *sql.DB
	var err error
	var connectionString string

	connectionString = fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, connectionName, dbName)
	if appengine.IsDevAppServer() {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, devConnection, dbName)
	}

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	return db
}
