package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/shijuvar/gokit/examples/http-api/model"
)

type NoteHandler struct {
	Repository model.Repository
}

func (h *NoteHandler) Post(c *gin.Context) {
	var note model.Note
	if err := c.Bind(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create note
	if _, err := h.Repository.Create(note); err != nil {
		if errors.Is(err, model.ErrNoteExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *NoteHandler) GetAll(c *gin.Context) {
	if notes, err := h.Repository.GetAll(); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, notes)
	}
}

func (h *NoteHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if note, err := h.Repository.GetById(id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, note)
	}
}
