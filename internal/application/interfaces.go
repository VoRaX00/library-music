package application

import "library-music/internal/domain"

type IMusicService interface {
	Add(music domain.MusicToAdd) (int, error)
	Delete(id int) error
	Update(music domain.MusicToUpdate, id int) error
	GetAll(params domain.MusicFilterParams, page int) ([]domain.MusicToGet, error)
	Get(song, group string) (domain.MusicToGet, error)
	GetText(song, group string, page int) (string, error)
}
