package main

import (
	"log"

	"github.com/gin-gonic/gin"

	apphttp "github.com/shijuvar/gokit/examples/http-api/http/gin"
	"github.com/shijuvar/gokit/examples/http-api/memstore"
)

func main() {
	repo, err := memstore.NewInmemoryRepository() // With in-memory database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.NoteHandler{
		Repository: repo, // Injecting dependency
	}
	router := gin.Default()
	// Routes
	router.GET("/api/notes", h.GetAll)
	router.GET("/api/notes/:id", h.Get)
	router.POST("/api/notes", h.Post)
	//router.PUT("/api/notes/:id", h.Put)
	//router.DELETE("/api/notes/:id", h.Delete)
	router.Run("localhost:8080")
}
