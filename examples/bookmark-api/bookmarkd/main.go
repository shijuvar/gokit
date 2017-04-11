package main

import (
	"log"
	"net/http"

	"github.com/shijuvar/gokit/examples/bookmark-api/routers"
	"github.com/shijuvar/gokit/examples/bookmark-api/bootstrapper"
)

// Entry point of the program
func main() {

	// Calls startup logic
	bootstrapper.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	// Create the Server
	server := &http.Server{
		Addr:    bootstrapper.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}
