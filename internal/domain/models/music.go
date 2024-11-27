package models

import "time"

type Music struct {
	Id          int       `json:"id" db:"id"`
	Song        string    `json:"song" db:"song"`
	Group       Group     `json:"group" db:"group"`
	Text        string    `json:"text" db:"text_song"`
	Link        string    `json:"link" db:"link" example:"https://example.com"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date" example:"DD.MM.YYYY"`
}
