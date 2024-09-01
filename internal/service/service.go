package service

import (
	"context"
	"kode-notes/internal/models"
)

type NoteService struct {
	s  NoteServicer
	sp Speller
}

func NewNoteService(s NoteServicer, sp Speller) *NoteService {
	return &NoteService{
		s:  s,
		sp: sp,
	}
}

type NoteServicer interface {
	Authorize(context.Context, *models.SignIn) (int, error)
	Authenticate(context.Context, int) (bool, error)
	AddNoteRepo(context.Context, *models.Note) error
	GetNotesListRepo(context.Context, int) ([]*models.Note, error)
}

type Speller interface {
	CheckText(string) ([]*models.Spell, error)
}
