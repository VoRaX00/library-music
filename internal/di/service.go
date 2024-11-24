package di

import (
	"library-music/internal/services/music"
	"library-music/internal/storage"
	"log/slog"
)

type Service struct {
	music.IMusicService
}

func NewService(log *slog.Logger, repos *storage.Repository) *Service {
	return &Service{
		IMusicService: music.New(log, repos.Music),
	}
}
