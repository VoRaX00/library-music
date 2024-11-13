package application

import (
	"library-music/internal/domain"
)

type IBaseRepository[T any, U any] interface {
	Add(T) (U, error)
	Delete(U) error
	Update(T, U) (T, error)
	GetById(U) (T, error)
}

type IMusicRepository interface {
	IBaseRepository[domain.Music, int]
	GetAll(params domain.Music, page int) ([]domain.Music, error)
	Get(song, group string) (domain.Music, error)
	GetText(song, group string) (string, error)
}
