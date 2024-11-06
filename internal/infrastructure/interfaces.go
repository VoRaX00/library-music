package infrastructure

import (
	"library-music/internal/domain"
)

type IMusicRepository interface {
	Add(music domain.Music) (int, error)
	Delete(id int) error
	Update(music domain.Music, id int) error
	GetAll(params domain.Music, page int) ([]domain.Music, error)
	Get(song, group string) (domain.Music, error)
	GetText(song, group string) (string, error)
}
