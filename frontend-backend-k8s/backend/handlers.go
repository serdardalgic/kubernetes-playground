package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	buffer.WriteString(`{"alive": true, `)

	err := Ping()
	if err != nil {
		buffer.WriteString(`"redis_conn": "shit"}`)
		fmt.Printf("err = %+v\n", err)
	} else {
		buffer.WriteString(`"redis_conn": "good"}`)
	}
	fmt.Fprintln(w, buffer.String())

}
func printHandler(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	var req_num int
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

	ok, err := Exists("requests")
	if ok {
		buffer.WriteString("\nRequests exists on redis\n")
	} else {
		buffer.WriteString("\nRedis integration is not implemented\n")
	}
	if r.URL.Path != "/favicon.ico" {
		Requests++
		req_num, _ = Incrby("requests", 1)
	}

	buffer.WriteString("\nTotal number of requests to this pod:")
	buffer.WriteString(strconv.Itoa(Requests))

	buffer.WriteString("\nTotal number of requests in all system:")
	buffer.WriteString(strconv.Itoa(req_num))

	buffer.WriteString("\nApp Uptime: ")
	t := time.Now()
	uptime := t.Sub(StartTime)
	buffer.WriteString(uptime.String())
	buffer.WriteString("\nLog Time: ")
	buffer.WriteString(t.String())
	fmt.Fprintln(w, buffer.String())
}
