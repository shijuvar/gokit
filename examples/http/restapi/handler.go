package main

import (
	// internal
	"encoding/json"
	"net/http"

	// external
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type handler struct {
	repository repository  // interface for persistence
	logger     *zap.Logger // Uber's Zap logger
}

//HTTP Post - /api/notes
func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.logger.Sync()
	var note note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		h.logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create note
	if err := h.repository.create(note); err != nil {
		h.logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("created note",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

//HTTP Get - /api/notes
func (h *handler) getAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	if notes, err := h.repository.getAll(); err != nil {
		h.logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == errNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		j, err := json.Marshal(notes)
		if err != nil {
			h.logger.Error(err.Error(),
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
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if note, err := h.repository.getById(id); err != nil {
		h.logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		if err == errNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(note)
		if err != nil {
			h.logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Put - /api/notes/{id}
func (h *handler) put(w http.ResponseWriter, r *http.Request) {
	defer h.logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	var note note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		h.logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	if err := h.repository.update(id, note); err != nil {
		h.logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("updated note",
		zap.String("note id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/notes/{id}
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	defer h.logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	if err := h.repository.delete(id); err != nil {
		h.logger.Error(err.Error(),
			zap.String("note id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("deleted note",
		zap.String("note id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
