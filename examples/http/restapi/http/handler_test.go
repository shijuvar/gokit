package http

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/shijuvar/gokit/examples/http/restapi/mocks"
	"github.com/shijuvar/gokit/examples/http/restapi/model"
)

var (
	r              *mux.Router
	mockCtrl       *gomock.Controller
	mockRepository *mocks.MockRepository
	logger         *zap.Logger
	handler        NoteHandler
	w              *httptest.ResponseRecorder
)

func tearTown() {
	mockRepository = nil
}

// setUp helper function set up before executing each unit tests
func setUp(t *testing.T) {
	r = mux.NewRouter()
	mockCtrl = gomock.NewController(t)
	mockRepository = mocks.NewMockRepository(mockCtrl)
	logger = zaptest.NewLogger(t)
	handler = NoteHandler{
		Repository: mockRepository,
		Logger:     logger,
	}
	w = httptest.NewRecorder()

}

func TestNoteHandler_Post_Valid_Note(t *testing.T) {
	setUp(t)
	mockRepository.EXPECT().Create(getMockNote()).Return(nil).Times(1)
	noteJson := `{"title": "mux", "description": "Gorilla mux is a router library"}`
	r.HandleFunc("/api/notes", handler.Post).Methods("POST")
	req, err := http.NewRequest(
		"POST",
		"/api/notes",
		strings.NewReader(noteJson),
	)
	if err != nil {
		t.Error(err)
	}

	r.ServeHTTP(w, req)
	assert.Equal(
		t,
		http.StatusCreated,
		w.Code,
		fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusCreated, w.Code),
	)
}

func TestNoteHandler_Post_Duplicate_Note_Title(t *testing.T) {
	setUp(t)
	mockRepository.EXPECT().Create(getMockNote()).Return(model.ErrNoteExists).Times(2)
	mockRepository.Create(getMockNote())
	noteJson := `{"title": "mux", "description": "Gorilla mux is a router library"}`
	r.HandleFunc("/api/notes", handler.Post).Methods("POST")
	req, err := http.NewRequest(
		"POST",
		"/api/notes",
		strings.NewReader(noteJson),
	)
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(
		t,
		http.StatusBadRequest,
		w.Code,
		fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusBadRequest, w.Code),
	)
}

func TestNoteHandler_GetAll(t *testing.T) {
	setUp(t)
	mockNotes := getMockNotes()
	mockRepository.EXPECT().GetAll().Return(mockNotes, nil)
	r.HandleFunc("/api/notes", handler.GetAll).Methods("GET")
	req, err := http.NewRequest(
		"GET",
		"/api/notes",
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(
		t,
		http.StatusOK,
		w.Code,
		fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusOK, w.Code),
	)
	var notes []model.Note
	json.Unmarshal(w.Body.Bytes(), &notes)
	assert.Equal(
		t,
		2,
		len(notes),
	)

	assert.Contains(
		t,
		notes,
		getMockNote(),
	)
}

func getMockNote() model.Note {
	return model.Note{
		Title:       "mux",
		Description: "Gorilla mux is a router library",
	}
}
func getMockNotes() []model.Note {
	notes := []model.Note{
		model.Note{
			Title:       "mux",
			Description: "Gorilla mux is a router library",
		},
		model.Note{
			Title:       "zap",
			Description: "Uber zap is a logging package",
		},
	}
	return notes
}
