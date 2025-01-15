package services

import (
	"library-music/internal/domain/models"
	"reflect"
)

type SongDetail struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type MusicToAdd struct {
	Song  string `json:"song"`
	Group string `json:"group"`
}

type MusicToUpdate struct {
	Song        string `json:"song,required" validate:"required"`
	Group       string `json:"group,required" validate:"required"`
	Text        string `json:"text,required" validate:"required"`
	Link        string `json:"link,required" validate:"required,url" example:"https://example.com"`
	ReleaseDate string `json:"releaseDate,required" db:"release_date" validate:"omitempty,datetime" example:"DD.MM.YYYY"`
}

type MusicToPartialUpdate struct {
	Song        string `json:"song,omitempty" validate:"omitempty"`
	Group       string `json:"group,omitempty" validate:"omitempty"`
	Text        string `json:"text,omitempty" validate:"omitempty"`
	Link        string `json:"link,omitempty" validate:"omitempty,url" example:"https://example.com"`
	ReleaseDate string `json:"releaseDate,omitempty" validate:"omitempty,datetime" example:"DD.MM.YYYY"`
}

func (m *MusicToPartialUpdate) ParsePartial() MusicToUpdate {
	var update MusicToUpdate
	partialVal := reflect.ValueOf(m).Elem()
	updateVal := reflect.ValueOf(&update).Elem()

	for i := 0; i < partialVal.NumField(); i++ {
		field := partialVal.Field(i)
		if !field.IsNil() {
			updateVal.FieldByName(partialVal.Type().Field(i).Name).Set(field.Elem())
		}
	}
	return update
}

type MusicToGet struct {
	Id          int          `json:"id"`
	Song        string       `json:"song"`
	Group       models.Group `json:"group"`
	Link        string       `json:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
	ReleaseDate string       `json:"releaseDate" example:"16.07.2006"`
}

type MusicFilterParams struct {
	Song        string `json:"song,omitempty" validate:"omitempty"`
	Group       string `json:"group,omitempty" validate:"omitempty"`
	Text        string `json:"text,omitempty" validate:"omitempty"`
	Link        string `json:"link,omitempty" validate:"omitempty,url" example:"https://example.com"`
	ReleaseDate string `json:"releaseDate" validate:"omitempty,datetime=02.01.2006" example:"DD.MM.YYYY"`
}

//func NewMusicFilterParams(song, group, text, link string, releaseDate time.Time) MusicFilterParams {
//	return MusicFilterParams{
//		Song:        song,
//		Group:       group,
//		Text:        text,
//		Link:        link,
//		ReleaseDate: releaseDate,
//	}
//}
