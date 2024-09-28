package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"library-music/internal/domain"
)

type MusicRepository struct {
	db *sqlx.DB
}

func NewMusicRepository(db *sqlx.DB) *MusicRepository {
	return &MusicRepository{
		db: db,
	}
}

func (r *MusicRepository) Add(music domain.Music) error {
	return nil
}

func (r *MusicRepository) Delete(song string) error {
	return nil
}

func (r *MusicRepository) Update(song string) error {
	return nil
}

func (r *MusicRepository) GetAll() ([]domain.Music, error) {
	return nil, nil
}

func (r *MusicRepository) Get(song string) (string, error) {
	return "", nil
}
