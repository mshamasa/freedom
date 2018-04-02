package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mshamasa/freedom/strength"
)

func main() {
	router := mux.NewRouter()

	router.Methods("GET").PathPrefix("/strength/{userID}").Handler(strength.IndexHandler)
	router.Methods("POST").PathPrefix("/strength/addRows").Handler(strength.AddRowsHandler)
	router.Methods("PUT").PathPrefix("/strength/saveWorkout").Handler(strength.SaveWorkoutHandler)
	router.Methods("PUT").PathPrefix("/strength/updateDate").Handler(strength.UpdateDateHandler)
	router.Methods("DELETE").PathPrefix("/strength/deleteRow").Handler(strength.DeleteRowHandler)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "application/json"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(router))
	http.ListenAndServe(":8080", router)
}
