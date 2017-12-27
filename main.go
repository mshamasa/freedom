package main

import (
	"log"
	"net/http"

	"github.com/mshamasa/freedom/strength"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Methods("GET").PathPrefix("/strength/{userID}").Handler(strength.IndexHandler)
	router.Methods("PUT").PathPrefix("/strength/addRows").Handler(strength.AddRowsHandler)
	router.Methods("PUT").PathPrefix("/strength/saveWorkout").Handler(strength.SaveWorkoutHandler)
	router.Methods("PUT").PathPrefix("/strength/updateDate").Handler(strength.UpdateDateHandler)
	router.Methods("DELETE").PathPrefix("/strength/deleteRow").Handler(strength.DeleteRowHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
