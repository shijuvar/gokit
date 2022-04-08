package http

import (
	// internal
	"encoding/json"
	"net/http"

	// external
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/shijuvar/gokit/examples/http/restapi/model"
)

type NoteHandler struct {
	Repository model.Repository // interface for persistence
	Logger     *zap.Logger      // Uber's Zap logger
}

//HTTP Post - /api/notes
func (h *NoteHandler) Post(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.Logger.Sync()
	var note model.Note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create note
	if err := h.Repository.Create(note); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == model.ErrNoteExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("created note",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

//HTTP Get - /api/notes
func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.Logger.Sync()
	// Get all
	if notes, err := h.Repository.GetAll(); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == model.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		j, err := json.Marshal(notes)
		if err != nil {
			h.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Get - /api/notes/{id}
func (h *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if note, err := h.Repository.GetById(id); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		if err == model.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(note)
		if err != nil {
			h.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Put - /api/notes/{id}
func (h *NoteHandler) Put(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	var note model.Note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	if err := h.Repository.Update(id, note); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("updated note",
		zap.String("note id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/notes/{id}
func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	defer h.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	if err := h.Repository.Delete(id); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("deleted note",
		zap.String("note id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
