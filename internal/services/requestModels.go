package services

import "time"

type MusicToAdd struct {
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Text        string `json:"text" db:"text_song"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}

type MusicToUpdate struct {
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Text        string `json:"text" db:"text_song"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}

type MusicToDelete struct {
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
}

type MusicToGet struct {
	Id          int    `json:"id" db:"id"`
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
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
