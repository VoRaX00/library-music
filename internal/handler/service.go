package handler

import (
	"library-music/internal/domain/models"
	"library-music/internal/services"
	"library-music/internal/services/externalApi"
	"library-music/internal/services/music"
	"library-music/internal/storage"
	"log/slog"
)

type Music interface {
	Add(music models.Music) (int, error)
	Delete(id int) error
	Update(music services.MusicToUpdate, id int) error
	GetAll(params services.MusicFilterParams, countSongs, page int) ([]services.MusicToGet, error)
	Get(song, group string) (services.MusicToGet, error)
	GetText(song, group string, countVerse, page int) (string, error)
}

type ExternalApi interface {
	Info(song, group string) (services.SongDetail, error)
}

type Service struct {
	Music       Music
	ExternalApi ExternalApi
}

func NewService(log *slog.Logger, repos *storage.Repository) *Service {
	return &Service{
		Music:       music.New(log, repos.Music),
		ExternalApi: externalApi.New(log),
	}
}
