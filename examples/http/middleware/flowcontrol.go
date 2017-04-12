package main

import (
	"fmt"
	"log"
	"net/http"
)

func middlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareFirst - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("MiddlewareFirst - After Handler")
	})
}

func middlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareSecond - Before Handler")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "pass123" {
				log.Println("Authorized to the system")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failed to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}

		log.Println("MiddlewareSecond - After Handler")
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index Handler")
	fmt.Fprintf(w, "Welcome")
}
func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing message Handler")
	fmt.Fprintf(w, "HTTP Middleware is awesome")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {

	http.HandleFunc("/favicon.ico", iconHandler)
	http.Handle("/", middlewareFirst(middlewareSecond(http.HandlerFunc(index))))
	http.Handle("/message", middlewareFirst(middlewareSecond(http.HandlerFunc(message))))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
