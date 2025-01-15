package music

import (
	"library-music/internal/domain/models"
)

type Repo interface {
	Add(music models.Music) (int, error)
	Delete(musicId int) error
	Update(music models.Music, id int) error
	GetById(musicId int) (models.Music, error)
	GetAll(params models.Music, countSongs, page int) ([]models.Music, error)
	Get(song, group string) (models.Music, error)
	GetText(song, group string) (string, error)
}
