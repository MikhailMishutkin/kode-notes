package transport

import (
	"context"
	"encoding/json"
	"io"
	"kode-notes/internal/models"
	"log"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

// ...
func (h *NoteHandle) AddNoteHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("AddNoteHandle was invoked")
	headerValue := http.Header.Get(r.Header, authorizationHeader)

	if headerValue == "" {
		http.Error(w, "empty auth header", http.StatusUnauthorized)
		return
	}

	userId, err := h.h.Authenticate(context.Background(), headerValue)
	if err != nil {
		http.Error(w, "wrong auth header", http.StatusUnauthorized)
		return

	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var note *models.Note
	err = json.Unmarshal(content, &note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	note.UserId = userId

	err = h.h.AddNote(context.Background(), note)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

// ...
func (h *NoteHandle) GetNotesListHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("GetNotesListHandle was invoked")
	headerValue := http.Header.Get(r.Header, authorizationHeader)
	if headerValue == "" {
		http.Error(w, "empty auth header", http.StatusUnauthorized)
		return
	}

	userId, err := h.h.Authenticate(context.Background(), headerValue)
	if err != nil {
		http.Error(w, "wrong auth header", http.StatusUnauthorized)
		return

	}

	response, err := h.h.GetNotesList(context.Background(), userId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}

	}
}
