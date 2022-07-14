package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apphttp "github.com/shijuvar/gokit/examples/http-echo/http"
	"github.com/shijuvar/gokit/examples/http-echo/memstore"
)

func main() {
	repo, err := memstore.NewInmemoryRepository() // With in-memory database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.NoteHandler{
		Repository: repo, // Injecting dependency
	}
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	// Routes
	e.GET("/api/notes", h.GetAll)
	e.GET("/api/notes/:id", h.Get)
	e.POST("/api/notes", h.Post)
	e.PUT("/api/notes/:id", h.Put)
	e.DELETE("/api/notes/:id", h.Delete)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
