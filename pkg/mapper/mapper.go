package mapper

import (
	"library-music/internal/domain/models"
	"library-music/internal/services"
	"time"
)

type MusicMapper struct {
}

func (m *MusicMapper) FilterToMusic(object services.MusicFilterParams) models.Music {
	date, _ := time.Parse("02.01.2006", object.ReleaseDate)
	return models.Music{
		Song:        object.Song,
		Group:       models.Group{Name: object.Group},
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}
}

func (m *MusicMapper) MusicForGet(object models.Music) services.MusicToGet {
	return services.MusicToGet{
		Id:          object.Id,
		Song:        object.Song,
		Group:       object.Group,
		Link:        object.Link,
		ReleaseDate: object.ReleaseDate.Format("02.01.2006"),
	}
}

func (m *MusicMapper) UpdateToMusic(object services.MusicToUpdate) (models.Music, error) {
	date, err := time.Parse("02.01.2006", object.ReleaseDate)
	if err != nil && object.ReleaseDate != "" {
		return models.Music{}, err
	}
	return models.Music{
		Song:        object.Song,
		Group:       models.Group{Name: object.Group},
		Text:        object.Text,
		Link:        object.Link,
		ReleaseDate: date,
	}, nil
}
