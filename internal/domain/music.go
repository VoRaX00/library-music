package domain

type Music struct {
	Id          int    `json:"id" db:"id"`
	Song        string `json:"song" db:"song"`
	Group       string `json:"group" db:"music_group"`
	Text        string `json:"text" db:"text_song"`
	Link        string `json:"link" db:"link"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
}
