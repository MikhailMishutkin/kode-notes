package service

import (
	"context"
	"errors"
	"fmt"
	"kode-notes/internal/models"
	"log"
	"strconv"
)

// ...
func (s *NoteService) Authorize(ctx context.Context, userData *models.SignIn) (id int, err error) {
	log.Println("Authorize service was invoked")
	id, err = s.s.Authorize(ctx, userData)
	if err != nil {
		return 0, fmt.Errorf("wrong username or password: %v", err)
	}

	return id, err
}

// ...
func (s *NoteService) Authenticate(ctx context.Context, auth string) (int, error) {
	log.Println("Authenticate service was invoked")
	authInt, err := strconv.Atoi(auth)

	if err != nil {
		return 0, fmt.Errorf("authorization string must be an integer: %v", err)
	}
	exist, err := s.s.Authenticate(ctx, authInt)
	if err != nil {

		return 0, fmt.Errorf("error to try find entry in db: %v", err)
	}
	if !exist {

		return 0, errors.New("authentication string is invalid")
	}

	return authInt, err
}
