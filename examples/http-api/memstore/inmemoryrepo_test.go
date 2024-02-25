package memstore

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/shijuvar/gokit/examples/http-api/model"
)

func TestNoteRepository_CreateValidNote(t *testing.T) {
	var wantError error
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	wantError = nil
	repository, _ := NewInmemoryRepository()
	_, err := repository.Create(note)
	assert.Equal(t, wantError, err, "failed to create Note")
}
func TestNoteRepository_CreateDuplicateNote(t *testing.T) {
	var wantError error
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	wantError = model.ErrNoteExists
	repository, _ := NewInmemoryRepository()
	_, err := repository.Create(note)
	_, err = repository.Create(note)
	assert.Equal(t, wantError, err, "failed to validate Note")
}

func TestNoteRepository_Create(t *testing.T) {

	repository, _ := NewInmemoryRepository()
	tests := []struct {
		name    string
		note    model.Note
		wantErr error
	}{
		{
			name: "Add Note with valid input",
			note: model.Note{
				Title:       "slog",
				Description: "slog is a logging package",
			},
			wantErr: nil,
		},
		{
			name: "Provide a duplicate Note",
			note: model.Note{
				Title:       "slog",
				Description: "slog is a logging package",
			},
			wantErr: model.ErrNoteExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := repository.Create(tt.note)
			assert.Equalf(t, tt.wantErr, err, "Create() error = %v, wantErr %v", err, tt.wantErr)
			if err == nil {
				assert.NotEmpty(t, id, "Create() error = NoteID is empty")

			}
		})
	}
}

func TestNoteRepository_Delete(t *testing.T) {
	repository, _ := NewInmemoryRepository()
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	id, _ := repository.Create(note)
	tests := []struct {
		name    string
		noteID  string
		wantErr error
	}{
		{
			name:    "delete Note",
			noteID:  id,
			wantErr: nil,
		},
		{
			name:    "delete same Note again",
			noteID:  id,
			wantErr: model.ErrNoteNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.Delete(tt.noteID)
			assert.Equalf(t, tt.wantErr, err, "Delete() error = %v, wantErr %v", err, tt.wantErr)

		})
	}
}

func TestCustomerRepository_GetAll(t *testing.T) {
	repository, _ := NewInmemoryRepository()
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	repository.Create(note)
	tests := []struct {
		name    string
		want    []model.Note
		wantErr error
	}{
		{
			name: "getting all Note",
			want: []model.Note{
				model.Note{
					Title:       "slog",
					Description: "slog is a logging package",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.GetAll()
			assert.Equalf(t, tt.wantErr, err, "GetAll() error = %v, wantErr %v", err, tt.wantErr)
			if err == nil {
				assert.Equalf(t, len(tt.want), len(got), "GetAll() got = %v, want %v", got, tt.want)
				assert.Equalf(t, tt.want[0].Title, got[0].Title, "GetAll() Note Title, got = %v, want %v", got[0].Title, tt.want[0].Title)

			}
		})
	}
}

func TestNoteRepository_GetById(t *testing.T) {
	repository, _ := NewInmemoryRepository()
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	id, _ := repository.Create(note)
	uid, _ := uuid.NewV4()
	tests := []struct {
		name    string
		id      string
		want    model.Note
		wantErr error
	}{
		{
			name: "getting by NoteID",
			id:   id,
			want: model.Note{
				Title:       "slog",
				Description: "slog is a logging package",
			},

			wantErr: nil,
		},
		{
			name:    "getting by non existing  NoteID",
			id:      uid.String(),
			want:    model.Note{},
			wantErr: model.ErrNoteNotExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.GetById(tt.id)
			assert.Equalf(t, tt.wantErr, err, "GetById() error = %v, wantErr %v", err, tt.wantErr)
			if err == nil {
				assert.Equalf(t, tt.want.Title, got.Title, "GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerRepository_Update(t *testing.T) {
	repository, _ := NewInmemoryRepository()
	note := model.Note{
		Title:       "slog",
		Description: "slog is a logging package",
	}
	id, _ := repository.Create(note)
	uid, _ := uuid.NewV4()
	tests := []struct {
		name    string
		id      string
		note    model.Note
		wantErr error
	}{
		{
			name: "getting by NoteID",
			id:   id,
			note: model.Note{
				Title:       "slog",
				Description: "slog is a standard library package for logging ",
			},

			wantErr: nil,
		},
		{
			name:    "updating non existing  NoteID",
			id:      uid.String(),
			wantErr: model.ErrNoteNotExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.Update(tt.id, tt.note)
			assert.Equalf(t, tt.wantErr, err, "Update() error = %v, wantErr %v", err, tt.wantErr)

		})
	}
}

func TestNewNoteRepository(t *testing.T) {
	repository, _ := NewInmemoryRepository()

	tests := []struct {
		name string
		want model.Repository
	}{
		{
			name: "check creation of Note repository",
			want: repository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewInmemoryRepository()
			assert.Equalf(t, tt.want, got, "NewNoteRepository() = %v, want %v", got, tt.want)
		})
	}
}
