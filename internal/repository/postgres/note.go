package postgres

import (
	"context"
	"fmt"
	"kode-notes/internal/models"
	"log"
)

// ...
func (r *Repo) AddNoteRepo(ctx context.Context, note *models.Note) error {
	log.Println("AddNoteRepo was invoked")

	const query = `INSERT INTO notes (title, note, user_id) VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(ctx, query, note.Title, note.Note, note.UserId)
	if err != nil {
		return fmt.Errorf("error to add note to db: %v", err)
	}

	return err
}

// ...
func (r *Repo) GetNotesListRepo(ctx context.Context, userId int) ([]*models.Note, error) {
	log.Println("GetNotesListRepo was invoked")

	const query = `SELECT title, note FROM notes WHERE user_id = $1`

	list := make([]*models.Note, 0, 10)

	rows, err := r.DB.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("error to add note to db: %v", err)
	}
	for rows.Next() {
		note := &models.Note{}
		if err = rows.Scan(&note.Title, &note.Note); err != nil {
			return nil, fmt.Errorf("some trouble with get list of notes: %v", err)
		}
		list = append(list, note)
	}

	return list, err
}
