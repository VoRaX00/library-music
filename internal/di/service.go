package di

import (
	"library-music/internal/domain/models"
	"library-music/internal/services"
	"library-music/internal/services/music"
	"library-music/internal/storage"
	"log/slog"
)

type Music interface {
	Add(music services.ToAdd) (int, error)
	Delete(id int) error
	Update(music services.ToUpdate, id int) (models.Music, error)
	GetAll(params services.FilterParams, page int) ([]services.ToGet, error)
	Get(song, group string) (services.ToGet, error)
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
