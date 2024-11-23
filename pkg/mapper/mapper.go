package mapper

import (
	"library-music/internal/domain"
	"library-music/internal/services"
	"time"
)

type MusicMapper struct {
}

func NewMapper[T any, U any]() *MusicMapper {
	return &MusicMapper{}
}

func (m *MusicMapper) FilterToMusic(object services.MusicFilterParams) domain.Music {
	return domain.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: object.ReleaseDate,
	}
}

func (m *MusicMapper) MusicForGet(object domain.Music) services.MusicToGet {
	return services.MusicToGet{
		Song:        object.Song,
		Group:       object.Group,
		Link:        object.Link,
		ReleaseDate: object.ReleaseDate.String(),
	}
}

func (m *MusicMapper) UpdateToMusic(object services.MusicToUpdate) (domain.Music, error) {
	date, err := time.Parse(object.ReleaseDate, "02-01-2006")
	if err != nil {
		return domain.Music{}, err
	}
	return domain.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}, nil
}

func (m *MusicMapper) AddToMusic(object services.MusicToAdd) (domain.Music, error) {
	date, err := time.Parse(object.ReleaseDate, "02-01-2006")
	if err != nil {
		return domain.Music{}, err
	}
	return domain.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}, nil
}
