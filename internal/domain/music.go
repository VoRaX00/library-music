package domain

type Music struct {
	Id    string `json:"id" db:"id"`
	Group string `json:"group" db:"music_group"`
	Song  string `json:"song" db:"song"`
	Text  string `json:"text" db:"text_song"`
	Link  string `json:"link" db:"link"`
}
