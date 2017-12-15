package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	util "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
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
		if err := server.ListenAndServe(); err != nil {
			util.Error.Fatalf("Error on starting the HTTP server:v%", err)
		}
		util.Info.Print("HTTP Server is started")

	}()

	interruptSignal := <-interrupt
	switch interruptSignal {
	case os.Kill:
		util.Error.Print("Got SIGKILL...")
	case os.Interrupt:
		util.Error.Print("Got SIGINT...")
	case syscall.SIGTERM:
		util.Error.Print("Got SIGTERM...")
	}

	util.Error.Print("The service is shutting down...")
	server.Shutdown(context.Background())
	util.Error.Print("Shut down is done")
}
