package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// VERSION Version of the program
const VERSION = "0.2"

// Requests Counter for number of Requests
var Requests int

// StartTime Time variable set before running Server
var StartTime time.Time

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
func printHandler(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	hostname, err := os.Hostname()
	if err != nil {
		log.Print(err, "Can not get Hostname")
	}
	buffer.WriteString("Hello from Kubernetes! Running on ")
	buffer.WriteString(hostname)
	buffer.WriteString(" | version: ")
	buffer.WriteString(VERSION)
	buffer.WriteString("\nTotal number of requests to this pod:")
	if r.URL.Path != "/favicon.ico" {
		Requests++
	}
	buffer.WriteString(strconv.Itoa(Requests))
	buffer.WriteString("\nApp Uptime: ")
	t := time.Now()
	uptime := t.Sub(StartTime)
	buffer.WriteString(uptime.String())
	buffer.WriteString("\nLog Time: ")
	buffer.WriteString(t.String())
	fmt.Fprintln(w, buffer.String())
}

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
