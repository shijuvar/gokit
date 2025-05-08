package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
)

type chain []func(http.Handler) http.Handler

func (c chain) thenFunc(h http.HandlerFunc) http.Handler {
	return c.then(h)
}
func (c chain) then(h http.Handler) http.Handler {
	for _, mw := range slices.Backward(c) {
		h = mw(h)
	}
	return h
}

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
		next.ServeHTTP(w, r)
		log.Println("MiddlewareSecond - After Handler")
	})
}
func admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Admin - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("Admin - After Handler")
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

func main() {
	mux := http.NewServeMux()
	base := chain{middlewareFirst, middlewareSecond}
	admin := append(base, admin)
	mux.Handle("/", base.thenFunc(index))
	mux.Handle("/message", admin.thenFunc(message))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
