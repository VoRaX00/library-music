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

func (r *MusicRepository) Add(music domain.Music) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}

	var musicId int
	query := "INSERT INTO music (music_group, song, text_song, link) values ($1, $2, $3, $4) RETURNING id"

	row := tx.QueryRow(query, music.Group, music.Song, music.Text, music.Link)
	err = row.Scan(&musicId)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	return musicId, tx.Commit()
}

func (r *MusicRepository) Delete(song string) error {
	query := "DELETE FROM music WHERE song=$1"
	_, err := r.db.Exec(query, song)
	return err
}

func (r *MusicRepository) Update(music domain.Music) error {
	return nil
}

func (r *MusicRepository) GetAll() ([]domain.Music, error) {
	var musics []domain.Music
	query := "SELECT music_group, song, text_song, link FROM music"
	if err := r.db.Select(&musics, query); err != nil {
		return nil, err
	}
	return musics, nil
}

func (r *MusicRepository) Get(song string) (domain.Music, error) {
	var music domain.Music
	query := "SELECT music_group, song, text_song, link FROM music WHERE song=$1"
	err := r.db.Get(&music, query, song)
	return music, err
}
