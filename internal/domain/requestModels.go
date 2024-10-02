package domain

import "time"

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
