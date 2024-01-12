package main

import (
	"log"
	"net/http"
)

const addr = "localhost:8080"

func mainHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello"))
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
