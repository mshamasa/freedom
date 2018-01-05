package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mshamasa/freedom/strength"
	"google.golang.org/appengine"
)

func main() {
	router := mux.NewRouter()

	router.Methods("GET").PathPrefix("/strength/{userID}").Handler(strength.IndexHandler)
	router.Methods("PUT").PathPrefix("/strength/addRows").Handler(strength.AddRowsHandler)
	router.Methods("PUT").PathPrefix("/strength/saveWorkout").Handler(strength.SaveWorkoutHandler)
	router.Methods("PUT").PathPrefix("/strength/updateDate").Handler(strength.UpdateDateHandler)
	router.Methods("DELETE").PathPrefix("/strength/deleteRow").Handler(strength.DeleteRowHandler)

	http.Handle("/", router)
	appengine.Main()

	// log.Fatal(http.ListenAndServe(":8080", router))
}
