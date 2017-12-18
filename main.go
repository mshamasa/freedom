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
  router.Methods("POST").PathPrefix("/strength/save").Handler(strength.SaveHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}
