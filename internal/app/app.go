package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"library-music/internal/app/server"
	"library-music/internal/di"
	"library-music/internal/handler"
	"library-music/internal/storage"
	"library-music/internal/storage/postgres"
	"log/slog"
)

type App struct {
	Server *server.Server
}

func New(log *slog.Logger, storagePath string, port string) *App {
	db, err := connectDB(storagePath)
	if err != nil {
		log.Warn(err.Error())
	}

	defer func() {
		_ = db.Close()
	}()

	repos := storage.NewRepository(db)
	srs := di.NewService(log, repos)
	handlers := handler.NewHandler(log, srs)

	srv := server.New(log, port, handlers.InitRouter())
	return &App{
		Server: srv,
	}
}

func connectDB(storagePath string) (*sqlx.DB, error) {
	//fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.SSLMode)
	db, err := postgres.New(storagePath)
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	if err = goose.Up(db.DB, "./storage/migrations"); err != nil {
		return db, fmt.Errorf("error upgrading database: %v", err)
	}
	return db, nil
}
