package application

import "library-music/internal/domain"

type IMusicService interface {
	Add(music MusicToAdd) (int, error)
	Delete(id int) error
	Update(music MusicToUpdate, id int) (domain.Music, error)
	GetAll(params MusicFilterParams, page int) ([]MusicToGet, error)
	Get(song, group string) (MusicToGet, error)
	GetText(song, group string, page int) (string, error)
}
