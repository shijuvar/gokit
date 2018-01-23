package main

import (
	"fmt"
	"log"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func middlewareFirst() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("MiddlewareFirst - Before Handler")
			h.ServeHTTP(w, r)
			log.Println("MiddlewareFirst - After Handler")
		})
	}
}

func middlewareSecond() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("MiddlewareSecond - Before Handler")
			if r.URL.Path == "/message" {
				if r.URL.Query().Get("password") == "pass123" {
					log.Println("Authorized to the system")
					h.ServeHTTP(w, r)
				} else {
					log.Println("Failed to authorize to the system")
					return
				}
			} else {
				h.ServeHTTP(w, r)
			}

			log.Println("MiddlewareSecond - After Handler")
		})
	}
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
	http.Handle("/", Adapt(http.HandlerFunc(index), middlewareSecond(), middlewareFirst()))
	http.Handle("/message", Adapt(http.HandlerFunc(message), middlewareSecond(), middlewareFirst()))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
