package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"library-music/internal/application"
)

type Repository struct {
	application.IMusicRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IMusicRepository: NewMusicRepository(db),
	}
}
