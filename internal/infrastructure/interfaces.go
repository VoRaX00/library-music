package infrastructure

import "library-music/internal/domain"

type IMusicRepository interface {
	Add(music domain.Music) error
	Delete(song string) error
	Update(song string) error
	GetAll() ([]domain.Music, error)
	Get(song string) (domain.Music, error)
}
