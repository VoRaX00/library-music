package services

import (
	"library-music/internal/domain/models"
	"time"
)

type MusicToAdd struct {
	Song        string `json:"song" db:"song" validate:"required"`
	Group       string `json:"group" db:"music_group" validate:"required"`
	Text        string `json:"text,omitempty" db:"text_song" validate:"omitempty"`
	Link        string `json:"link" db:"link" validate:"required,url"`
	ReleaseDate string `json:"releaseDate" db:"release_date" validate:"required,datetime=02-01-2006"`
}

type MusicToUpdate struct {
	Song        string `json:"song,required" db:"song" validate:"required"`
	Text        string `json:"text,required" db:"text_song" validate:"required"`
	Link        string `json:"link,required" db:"link" validate:"required,url"`
	ReleaseDate string `json:"releaseDate,required" db:"release_date" validate:"omitempty,datetime=02-01-2006"`
}

type MusicToPartialUpdate struct {
	Song        *string `json:"song,omitempty" db:"song" validate:"omitempty"`
	Text        *string `json:"text,omitempty" db:"text_song" validate:"omitempty"`
	Link        *string `json:"link,omitempty" db:"link" validate:"omitempty,url"`
	ReleaseDate *string `json:"releaseDate,omitempty" db:"release_date" validate:"omitempty,datetime=02-01-2006"`
}

type MusicToGet struct {
	Id          int          `json:"id" db:"id"`
	Song        string       `json:"song" db:"song"`
	Group       models.Group `json:"group" db:"music_group"`
	Link        string       `json:"link" db:"link"`
	ReleaseDate string       `json:"releaseDate" db:"release_date"`
}

type MusicFilterParams struct {
	Song        string    `json:"song"`
	Group       string    `json:"group"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"releaseDate"`
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
