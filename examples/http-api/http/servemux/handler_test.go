package servemux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/shijuvar/gokit/examples/http-api/mocks"
	"github.com/shijuvar/gokit/examples/http-api/model"
)

var (
	r              *http.ServeMux
	mockCtrl       *gomock.Controller
	mockRepository *mocks.MockRepository
	handler        NoteHandler
	w              *httptest.ResponseRecorder
)

func tearTown() {
	mockRepository = nil
}

// setUp helper function set up before executing each unit tests
func setUp(t *testing.T) {
	r = http.NewServeMux()
	mockCtrl = gomock.NewController(t)
	mockRepository = mocks.NewMockRepository(mockCtrl)
	handler = NoteHandler{
		Repository: mockRepository,
	}
	w = httptest.NewRecorder()

}

func TestNoteHandler_Post_Valid_Note(t *testing.T) {
	setUp(t)
	uid, _ := uuid.NewV4()

	mockRepository.EXPECT().Create(getMockNote()).Return(uid.String(), nil).Times(1)
	noteJson := `{"title": "mux", "description": "Gorilla mux is a router library"}`
	r.HandleFunc("POST /api/notes", handler.Post)
	req, err := http.NewRequest(
		"POST",
		"/api/notes",
		strings.NewReader(noteJson),
	)
	if err != nil {
		assert.Error(t, err)
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
	mockRepository.EXPECT().Create(getMockNote()).Return("", model.ErrNoteExists).Times(2)
	mockRepository.Create(getMockNote())
	noteJson := `{"title": "mux", "description": "Gorilla mux is a router library"}`
	r.HandleFunc("POST /api/notes", handler.Post)
	req, err := http.NewRequest(
		"POST",
		"/api/notes",
		strings.NewReader(noteJson),
	)
	if err != nil {
		assert.Error(t, err)
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
	r.HandleFunc("GET /api/notes", handler.GetAll)
	req, err := http.NewRequest(
		"GET",
		"/api/notes",
		nil,
	)
	if err != nil {
		assert.Error(t, err)
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
