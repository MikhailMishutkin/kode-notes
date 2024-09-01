package transport

import (
	"context"
	"kode-notes/internal/models"

	"github.com/gorilla/mux"
)

type NoteHandle struct {
	h NoteHandler
}

func NewNoteHandle(h NoteHandler) *NoteHandle {
	return &NoteHandle{h: h}
}

type NoteHandler interface {
	Authorize(context.Context, *models.SignIn) (int, error)
	Authenticate(context.Context, string) (int, error)
	AddNote(context.Context, *models.Note) error
	GetNotesList(context.Context, int) ([]*models.Note, error)
}

// ...
func (h *NoteHandle) RegisterAuth(router *mux.Router) {
	router.HandleFunc("/auth", h.Auth).Methods("POST")
}

// ...
func (h *NoteHandle) RegisterNotes(router *mux.Router) {
	router.HandleFunc("/note/add", h.AddNoteHandle).Methods("POST")
	router.HandleFunc("/note/list", h.GetNotesListHandle).Methods("GET")
}
