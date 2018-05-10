package main

import (
	"net/http"

	// "github.com/aws/aws-lambda-go/lambda"

	"github.com/gorilla/mux"
	"github.com/mshamasa/freedom/strength"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/strength/{userID}", strength.IndexHandler).Methods("GET")
	router.HandleFunc("/strength/deleteRow", strength.DeleteRowHandler).Methods("DELETE")
	router.HandleFunc("/strength/addRows", strength.AddRowsHandler).Methods("POST")
	router.HandleFunc("/strength/updateDate", strength.UpdateDateHandler).Methods("PUT")
	router.HandleFunc("/strength/saveWorkout", strength.SaveWorkoutHandler).Methods("PUT")

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "application/json"})
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	http.Handle("/", (router))
	http.ListenAndServe(":8080", nil)
	// lambda.Start(router)
}
