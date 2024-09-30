package infrastructure

import "library-music/internal/domain"

type IMusicRepository interface {
	Add(music domain.MusicToAdd) (int, error)
	Delete(music domain.MusicToDelete) error
	Update(music domain.MusicToUpdate) error
	GetAll() ([]domain.Music, error)
	Get(music domain.MusicToGet) (domain.Music, error)
}
