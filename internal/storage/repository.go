package storage

import (
	"github.com/jmoiron/sqlx"
	"library-music/internal/services"
)

type Repository struct {
	services.Music
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Music: NewMusicRepository(db),
	}
}
