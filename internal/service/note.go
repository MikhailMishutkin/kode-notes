package service

import (
	"context"
	"fmt"
	"kode-notes/internal/models"
	"log"
	"net/url"
	"slices"
)

// ...
func (s *NoteService) AddNote(ctx context.Context, note *models.Note) error {
	log.Println("AddNote service was invoked")
	errSpell, err := s.sp.CheckText(note.Title)
	if err != nil {
		return err
	}
	mapTitle, err := HandleSpellResponse(errSpell, note.Title)
	if err != nil {
		return err
	}

	errSpell, err = s.sp.CheckText(note.Note)
	if err != nil {
		return err
	}

	mapNote, err := HandleSpellResponse(errSpell, note.Note)
	if err != nil {
		return err
	}

	note.Title = FromMapToString(mapTitle)
	note.Note = FromMapToString(mapNote)

	err = s.s.AddNoteRepo(ctx, note)
	if err != nil {
		return err
	}
	return err
}

// ...
func (s *NoteService) GetNotesList(ctx context.Context, userId int) (notes []*models.Note, err error) {

	log.Println("GetNotesList service was invoked")

	notes, err = s.s.GetNotesListRepo(ctx, userId)
	return notes, err
}

// ...
func HandleSpellResponse(errSpell []*models.Spell, s string) (m map[int]string, err error) {
	log.Println("HandleSpellResponse service was invoked")

	mapOriginal := MappingOriginal(s)

	//change word with error to the first of offered from speller response
	for _, v := range errSpell {
		if v.Code != 0 {
			utf, err := url.QueryUnescape(v.S[0])
			if err != nil {
				return nil, fmt.Errorf("something wrong with unicode in speller response: %v", err)
			}
			mapOriginal[v.Pos] = utf
		}
	}
	return mapOriginal, nil
}

// ...
func MappingOriginal(note string) map[int]string {
	mapOriginal := make(map[int]string)
	noteR := []rune(note)
	for i, v := range noteR {
		var word []rune
		if string(v) == " " {
			for k := i + 1; k < len(noteR); k++ {

				if noteR[k] == v {
					break
				} else {
					word = append(word, noteR[k])

				}
			}
			mapOriginal[i+1] = string(word)

		}

	}

	var word []rune
	for i, v := range noteR {

		if string(v) == " " {
			break
		} else {
			word = append(word, noteR[i])
		}

	}
	mapOriginal[0] = string(word)

	return mapOriginal
}

// ...
func FromMapToString(m map[int]string) (s string) {
	nums := make([]int, 0)
	for i := range m {
		nums = append(nums, i)
	}
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		s = s + m[nums[i]]
		s = s + " "
	}

	return s
}
