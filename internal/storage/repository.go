package storage

import (
	"github.com/jmoiron/sqlx"
	"library-music/internal/storage/music"
)

type Repository struct {
	Music *musicrepo.Music
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Music: musicrepo.New(db),
	}
}
