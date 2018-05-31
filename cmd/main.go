package main

import (
    "log"
	"net/http"
	
	"github.com/hrafn/gophercon2018/pkg/routing"
)

func main() {
    log.Printf("Server Is Starting.")
    r := routing.BaseRouter()
	http.ListenAndServe(":8000", r)
}
