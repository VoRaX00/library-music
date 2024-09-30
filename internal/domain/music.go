package domain

type Music struct {
	Id    string `json:"id" db:"id"`
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
	Text  string `json:"text" db:"text_song"`
	Link  string `json:"link" db:"link"`
}

type MusicToAdd struct {
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
	Text  string `json:"text" db:"text_song"`
	Link  string `json:"link" db:"link"`
}

type MusicToUpdate struct {
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
	Text  string `json:"text" db:"text_song"`
	Link  string `json:"link" db:"link"`
}

type MusicToDelete struct {
	Song string `json:"song" db:"song"`
}

type MusicToGet struct {
	Song string `json:"song" db:"song"`
}
