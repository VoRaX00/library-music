package app

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"library-music/internal/app/server"
	"library-music/internal/handler"
	"library-music/internal/storage"
	"library-music/internal/storage/postgres"
	"log/slog"
)

type App struct {
	Server *server.Server
	DB     *sqlx.DB
}

func New(log *slog.Logger, storagePath string, port string) *App {
	db, err := connectDB(storagePath)
	if err != nil {
		log.Warn(err.Error())
	}

	repos := storage.NewRepository(db)
	srs := handler.NewService(log, repos)
	handlers := handler.NewHandler(srs)

	srv := server.New(log, port, handlers.InitRouter())
	return &App{
		Server: srv,
		DB:     db,
	}
}

func (a *App) Stop(ctx context.Context) {
	err := a.DB.Close()
	if err != nil {
		slog.Error(err.Error())
	}
	a.Server.Stop(ctx)
}

func connectDB(storagePath string) (*sqlx.DB, error) {
	db, err := postgres.New(storagePath)
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	if err = goose.Up(db.DB, "./storage/migrations"); err != nil {
		return db, fmt.Errorf("error upgrading database: %v", err)
	}
	return db, nil
}
