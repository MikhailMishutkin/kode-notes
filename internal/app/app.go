package app

import (
	"context"
	"fmt"
	"kode-notes/configs"
	"kode-notes/internal/repository/postgres"
	"kode-notes/internal/repository/speller"
	"kode-notes/internal/service"
	"kode-notes/internal/transport"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"log/slog"

	"github.com/gorilla/mux"
)

type Server struct {
	noteRouter *mux.Router
	Logger     *slog.Logger
}

func Start(conf configs.Config) error {
	s := &Server{
		noteRouter: mux.NewRouter(),
		Logger:     slog.Default(),
	}
	db, err := NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	//dependencies
	repoDB := postgres.NewRepo(db)
	repoSp := speller.NewSpell(http.DefaultClient)
	noteService := service.NewNoteService(repoDB, repoSp)
	noteHandler := transport.NewNoteHandle(noteService)
	noteHandler.RegisterNotes(s.noteRouter)
	noteHandler.RegisterAuth(s.noteRouter)

	log.Println("Starting MessageService at port: 8080")
	return http.ListenAndServe(":8080", s)

}

func NewDB() (*pgx.Conn, error) {
	c, err := configs.New("./configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("can't load config to db: %v", err)
	}

	psqlInfo := fmt.Sprint(c.DB.ConnSql)

	db, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v", err)
	}
	//run if only db in docker
	//
	// m, err := migrate.New(
	// 	"/migrations/",
	// 	"postgres://root:root@localhost:5444/kode_notes?sslmode=disable",
	// 	//+c.DB.Migrate,
	// )
	// if err != nil {
	// 	return db, fmt.Errorf("can't automigrate: %v", err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Println(err)
	// 	//TODO
	// 	fmt.Errorf("%v", err)
	// }
	return db, err
}

// ServeHTTP
func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.noteRouter.ServeHTTP(w, r)
}
