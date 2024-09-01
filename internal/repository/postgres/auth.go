package postgres

import (
	"context"
	"errors"
	"kode-notes/internal/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// ...
func (r *Repo) Authorize(ctx context.Context, data *models.SignIn) (id int, err error) {
	log.Println("Authorize repo was invoked")

	const query = `SELECT id FROM users WHERE name = $1 and pass = $2`

	err = r.DB.QueryRow(ctx, query, data.UserName, data.Pass).Scan(&id)

	return id, err
}

// ...
func (r *Repo) Authenticate(ctx context.Context, auth int) (bool, error) {
	log.Println("Authenticate repo was invoked")

	const query = `SELECT 1 FROM users WHERE id = $1`
	var n int64
	err := r.DB.QueryRow(ctx, query, auth).Scan(&n)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil

}
