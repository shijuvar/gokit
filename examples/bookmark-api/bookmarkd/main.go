package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shijuvar/gokit/examples/bookmark-api/bootstrapper"
	"github.com/shijuvar/gokit/examples/bookmark-api/routers"
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

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Print("The HTTP Server is ready to listen and serve.")

	killSignal := <-interrupt
	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL...")
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Print("Done")
}
