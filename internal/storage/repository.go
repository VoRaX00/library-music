package storage

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Music *Music
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Music: New(db),
	}
}
