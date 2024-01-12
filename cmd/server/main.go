package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

const addr = "localhost:8080"

type Action string

const (
	Update Action = "update"
)

type MetricType string

const (
	Gauge   MetricType = "gauge"
	Counter            = "counter"
)

func mainHandler(resp http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(parts) != 4 {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	var (
		action     = Action(parts[0])
		metricType = MetricType(parts[1])
		_          = parts[2]
		value      = parts[3]
	)

	switch Action(action) {
	case Update:
	default:
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	switch MetricType(metricType) {
	case Gauge:
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

	case Counter:
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

	default:
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	log.Print("start server")

	server := http.Server{Addr: addr, Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
