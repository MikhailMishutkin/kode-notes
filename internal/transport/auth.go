package transport

import (
	"context"
	"encoding/json"
	"io"
	"kode-notes/internal/models"
	"log"
	"net/http"
)

// ...
func (h *NoteHandle) Auth(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth transport was invoked")

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var input *models.SignIn
	err = json.Unmarshal(content, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	id, err := h.h.Authorize(context.Background(), input)
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
		if err := json.NewEncoder(w).Encode(id); err != nil {
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}

	}

}
