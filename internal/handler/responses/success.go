package responses

import (
	"library-music/internal/services"
)

type SuccessID struct {
	ID int `json:"id"`
}

type SuccessStatus struct {
	Status string `json:"status"`
}

type SuccessMusics struct {
	Music []services.MusicToGet `json:"songs"`
}

type SuccessText struct {
	Text string `json:"text"`
}
