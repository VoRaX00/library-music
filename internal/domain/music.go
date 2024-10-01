package domain

type Music struct {
	Id    int    `json:"id" db:"id"`
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
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
}

type MusicToGet struct {
	Id    int    `json:"id" db:"id"`
	Song  string `json:"song" db:"song"`
	Group string `json:"group" db:"music_group"`
	Link  string `json:"link" db:"link"`
}
