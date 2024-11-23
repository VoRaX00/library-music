package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"library-music/internal/config"
	"library-music/internal/handler"
	"library-music/internal/services"
	"library-music/internal/storage"
	"library-music/internal/storage/postgres"
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
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	db, err := postgres.New(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.SSLMode))

	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	defer func() {
		_ = db.Close()
	}()

	if err = goose.Up(db.DB, "./storage/migrations"); err != nil {
		logrus.Fatalf("Error upgrading database: %v", err)
	}

	repos := storage.NewRepository(db)
	services := services.NewService(log, repos)
	handlers := handler.NewHandler(log, services)

	server := new(Server)

	go func() {
		if err = server.Start(cfg.Server.Port, handlers.InitRouter()); err != nil {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	logrus.Info("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

//func initConfig() error {
//	viper.AddConfigPath("configs")
//	viper.SetConfigName("config")
//	return viper.ReadInConfig()
//}
//
//func initDBConfig() storage.Config {
//	return storage.Config{
//		Host:     viper.GetString("db.host"),
//		Port:     viper.GetString("db.port"),
//		Username: viper.GetString("db.username"),
//		Password: os.Getenv("DB_PASSWORD"),
//		Database: viper.GetString("db.name"),
//		SSLMode:  viper.GetString("db.sslmode"),
//	}
//}

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
