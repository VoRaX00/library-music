package music

import (
	"library-music/internal/domain"
	"library-music/internal/services"
)

type IMusicService interface {
	Add(music services.MusicToAdd) (int, error)
	Delete(id int) error
	Update(music services.MusicToUpdate, id int) (domain.Music, error)
	GetAll(params services.MusicFilterParams, page int) ([]services.MusicToGet, error)
	Get(song, group string) (services.MusicToGet, error)
	GetText(song, group string, page int) (string, error)
}
