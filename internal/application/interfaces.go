package application

import "library-music/internal/domain"

type IMusicService interface {
	Add(music domain.Music) (int, error)
	Delete(song string) error
	Update(music domain.Music) error
	GetAll() ([]domain.Music, error)
	Get(song string) (domain.Music, error)
	GetText(song string) (string, error)
}
