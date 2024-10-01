package domain

type MusicFilterParams struct {
	Song  string `json:"song"`
	Group string `json:"group"`
	Text  string `json:"text"`
	Link  string `json:"link"`
}

func NewMusicFilterParams(song, group, text, link string) MusicFilterParams {
	return MusicFilterParams{
		Song:  song,
		Group: group,
		Text:  text,
		Link:  link,
	}
}
