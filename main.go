package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/mshamasa/freedom/strength"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  router.Methods("GET").PathPrefix("/strength/{userId}").Handler(strength.IndexHandler)
  router.HandleFunc("/add", Add)
  router.HandleFunc("/edit/{id}", Edit)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func Add(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Edit(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  fmt.Fprintln(w, "Passing Id:", id)
}