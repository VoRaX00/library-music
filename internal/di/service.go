package di

import (
	"library-music/internal/domain"
	"library-music/internal/services/music"
	"library-music/internal/storage"
	"log/slog"
)

type Music interface {
	Add(music music.ToAdd) (int, error)
	Delete(id int) error
	Update(music music.ToUpdate, id int) (domain.Music, error)
	GetAll(params music.FilterParams, page int) ([]music.ToGet, error)
	Get(song, group string) (music.ToGet, error)
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
