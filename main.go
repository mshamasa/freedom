package main

import (
	"log"
	"net/http"

	"github.com/mshamasa/freedom/strength"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Methods("GET").PathPrefix("/strength/{userId}").Handler(strength.IndexHandler)
	router.Methods("POST").PathPrefix("/strength/save").Handler(strength.SaveRowHandler)
	router.Methods("PUT").PathPrefix("/strength/saveWorkout").Handler(strength.SaveWorkoutHandler)
	router.Methods("PUT").PathPrefix("/strength/updateDate").Handler(strength.UpdateDateHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
