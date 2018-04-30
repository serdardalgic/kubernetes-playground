package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// VERSION Version of the program
const VERSION = "0.3"

// Requests Counter for number of Requests
var Requests int

// StartTime Time variable set before running Server
var StartTime time.Time

func main() {
	StartTime = time.Now()

	r := mux.NewRouter()
	r.HandleFunc("/", printHandler)
	r.HandleFunc("/health_check", HealthCheckHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
