package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"library-music/internal/app"
	"library-music/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// @title Library music API
// @version 1.0
// @description API Server for Library music services
// @host localhost:8090
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	storagePath := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.SSLMode)

	application := app.New(log, storagePath, cfg.Server.Port)

	log.Info("starting server")
	go application.Server.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	log.Info("stopping server", slog.String("signal", sig.String()))

	application.Server.Stop(context.Background())
	log.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}
	return log
}
