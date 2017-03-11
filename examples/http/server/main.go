package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Programming")
}

func main() {

	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
