package httpecho

import (
	"github.com/labstack/echo/v4"
	"github.com/shijuvar/gokit/examples/http-echo/model"
	"net/http"
)

type NoteHandler struct {
	Repository model.Repository
}

func (h *NoteHandler) Post(c echo.Context) error {
	var note model.Note
	if err := c.Bind(&note); err != nil {
		return err
	}
	// Create note
	if err := h.Repository.Create(note); err != nil {
		if err == model.ErrNoteExists {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}
func (h *NoteHandler) GetAll(c echo.Context) error {
	if notes, err := h.Repository.GetAll(); err != nil {
		if err == model.ErrNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, notes)
	}
}

func (h *NoteHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if note, err := h.Repository.GetById(id); err != nil {
		if err == model.ErrNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, note)
	}
}

func (h *NoteHandler) Put(c echo.Context) error {
	id := c.Param("id")
	var note model.Note
	if err := c.Bind(&note); err != nil {
		return err
	}
	// Update
	if err := h.Repository.Update(id, note); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *NoteHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.Repository.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
