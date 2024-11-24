package music

import "time"

type ToAdd struct {
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Text        string `json:"text" db:"text_song"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}

type ToUpdate struct {
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Text        string `json:"text" db:"text_song"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}

type ToDelete struct {
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
}

type ToGet struct {
	Id          int    `json:"id" db:"id"`
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}

type FilterParams struct {
	Song        string    `json:"song"`
	Group       string    `json:"group"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"releaseDate"`
}

func NewMusicFilterParams(song, group, text, link string, releaseDate time.Time) FilterParams {
	return FilterParams{
		Song:        song,
		Group:       group,
		Text:        text,
		Link:        link,
		ReleaseDate: releaseDate,
	}
}
