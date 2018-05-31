package routing

import (
    "net/http"
    "github.com/gorilla/mux"
)

func DiagnosticsRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/healthyz", homeHandler()).Methods(http.MethodGet)
    r.HandleFunc("/readyz", homeHandler()).Methods(http.MethodGet)
    return r
}

