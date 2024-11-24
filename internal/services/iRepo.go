package services

import (
	"library-music/internal/domain"
)

type Music interface {
	Add(music domain.Music) (int, error)
	Delete(musicId int) error
	Update(music domain.Music, id int) (domain.Music, error)
	GetById(musicId int) (domain.Music, error)
	GetAll(params domain.Music, page int) ([]domain.Music, error)
	Get(song, group string) (domain.Music, error)
	GetText(song, group string) (string, error)
}