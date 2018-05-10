package strength

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// IndexHandler is the entry point to getting all user data.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	req := Request{UserID: userID}

	list := Index(req)

	response := Response{list, Workout{}, "", ""}

	EncodeResponse(w, response)
}

// DeleteRowHandler is the entry point to delete a row.
func DeleteRowHandler(w http.ResponseWriter, r *http.Request) {
	var row Row
	var request Request

	q := r.URL.Query()
	request.UserID = q["userID"][0]
	rowIds := q["rowIds"][0]
	rowIDList := strings.Split(rowIds, ",")
	if len(rowIDList) > 0 {
		for i := 0; i < len(rowIDList); i++ {
			if id, err := strconv.ParseInt(rowIDList[i], 10, 32); err == nil {
				row.RowIds = append(row.RowIds, int32(id))
			}
		}
		request.Row = row
	}

	DeleteRow(request)

	response := Response{nil, Workout{}, "", ""}

	EncodeResponse(w, response)
}

// AddRowsHandler is the entry point to adding rows.
func AddRowsHandler(w http.ResponseWriter, r *http.Request) {
	if request, err := DecodeStrengthBodyRequest(r); err == nil {
		AddRows(request)

		list := Index(request)

		response := Response{list, Workout{}, "", ""}

		EncodeResponse(w, response)
	} else {
		log.Panicf("%s Error: ", err)
	}

}

// UpdateDateHandler handles updating date for all rows passed
func UpdateDateHandler(w http.ResponseWriter, r *http.Request) {
	if request, err := DecodeStrengthBodyRequest(r); err == nil {
		UpdateRowsDate(request)

		response := Response{nil, Workout{}, "", ""}

		EncodeResponse(w, response)
	} else {
		log.Panicf("%s Error: ", err)
	}
}

// SaveWorkoutHandler is the entry point to updating a single row.
func SaveWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	if request, err := DecodeStrengthBodyRequest(r); err == nil {
		workout := SaveWorkout(request)

		response := Response{nil, workout, "", ""}

		EncodeResponse(w, response)
	} else {
		log.Panicf("%s Error: ", err)
	}
}
