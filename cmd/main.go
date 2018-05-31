package main

import (
    "net/http"
    "fmt"
    "log"

    "github.com/gorilla/mux"
)

func main() {
    log.Printf("Server Is Starting.")
    r := mux.NewRouter()
    r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
    http.ListenAndServe(":8000", r)
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request is processing %s", r.URL.Path)
        fmt.Fprint(w, "Hello! Your request was processed.")
    }
}