package handler

import (
	"library-music/internal/services"
	"library-music/internal/services/music"
	"library-music/internal/storage"
	"log/slog"
)

type Music interface {
	Add(music services.MusicToAdd) (int, error)
	Delete(id int) error
	Update(music services.MusicToUpdate, id int) error
	GetAll(params services.MusicFilterParams, page int) ([]services.MusicToGet, error)
	Get(song, group string) (services.MusicInfo, error)
	GetText(song, group string, page int) (string, error)
}

type Service struct {
	Music Music
}

func NewService(log *slog.Logger, repos *storage.Repository) *Service {
	return &Service{
		Music: music.New(log, repos.Music),
	}
}
