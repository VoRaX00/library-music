package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"library-music/internal/application"
	"library-music/internal/handler"
	"library-music/internal/infrastructure"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}

	cfg := initDBConfig()

	db, err := infrastructure.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	defer func() {
		_ = db.Close()
	}()

	repos := infrastructure.NewRepository(db)
	services := application.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(Server)

	go func() {
		if err = server.Start(viper.GetString("server.port"), handlers.InitRouter()); err != nil {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	logrus.Info("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDBConfig() infrastructure.Config {
	return infrastructure.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}
