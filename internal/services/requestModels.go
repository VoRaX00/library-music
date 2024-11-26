package services

import (
	"library-music/internal/domain/models"
	"reflect"
	"time"
)

type MusicToAdd struct {
	Song        string `json:"song" db:"song" validate:"required"`
	Group       string `json:"group" db:"group" validate:"required"`
	Text        string `json:"text,omitempty" db:"text_song" validate:"omitempty"`
	Link        string `json:"link" db:"link" validate:"required,url"`
	ReleaseDate string `json:"releaseDate" db:"release_date" validate:"required,datetime" example:"DD-MM-YYYY"`
}

type MusicToUpdate struct {
	Song        string `json:"song,required" db:"song" validate:"required"`
	Text        string `json:"text,required" db:"text_song" validate:"required"`
	Link        string `json:"link,required" db:"link" validate:"required,url"`
	ReleaseDate string `json:"releaseDate,required" db:"release_date" validate:"omitempty,datetime" example:"DD-MM-YYYY"`
}

type MusicToPartialUpdate struct {
	Song        *string `json:"song,omitempty" db:"song" validate:"omitempty"`
	Text        *string `json:"text,omitempty" db:"text_song" validate:"omitempty"`
	Link        *string `json:"link,omitempty" db:"link" validate:"omitempty,url"`
	ReleaseDate *string `json:"releaseDate,omitempty" db:"release_date" validate:"omitempty,datetime" example:"DD-MM-YYYY"`
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
	Id          int          `json:"id" db:"id"`
	Song        string       `json:"song" db:"song"`
	Group       models.Group `json:"group" db:"group"`
	Link        string       `json:"link" db:"link"`
	ReleaseDate string       `json:"releaseDate" db:"release_date"`
}

type MusicFilterParams struct {
	Song        string    `json:"song"`
	Group       string    `json:"group"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"releaseDate" example:"DD-MM-YYYY"`
}

func NewMusicFilterParams(song, group, text, link string, releaseDate time.Time) MusicFilterParams {
	return MusicFilterParams{
		Song:        song,
		Group:       group,
		Text:        text,
		Link:        link,
		ReleaseDate: releaseDate,
	}
}
