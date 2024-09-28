package infrastructure

import "github.com/jmoiron/sqlx"

type Repository struct {
	IMusicRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IMusicRepository: NewMusicRepository(db),
	}
}
