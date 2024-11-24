package mapper

import (
	"library-music/internal/domain/models"
	"library-music/internal/services"
	"time"
)

type MusicMapper struct {
}

func (m *MusicMapper) FilterToMusic(object services.FilterParams) models.Music {
	return models.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: object.ReleaseDate,
	}
}

func (m *MusicMapper) MusicForGet(object models.Music) services.ToGet {
	return services.ToGet{
		Song:        object.Song,
		Group:       object.Group,
		Link:        object.Link,
		ReleaseDate: object.ReleaseDate.String(),
	}
}

func (m *MusicMapper) UpdateToMusic(object services.ToUpdate) (models.Music, error) {
	date, err := time.Parse(object.ReleaseDate, "02-01-2006")
	if err != nil {
		return models.Music{}, err
	}
	return models.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}, nil
}

func (m *MusicMapper) AddToMusic(object services.ToAdd) (models.Music, error) {
	date, err := time.Parse(object.ReleaseDate, "02-01-2006")
	if err != nil {
		return models.Music{}, err
	}
	return models.Music{
		Song:        object.Song,
		Group:       object.Group,
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}, nil
}
