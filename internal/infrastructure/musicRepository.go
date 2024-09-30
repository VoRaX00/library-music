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

func (r *MusicRepository) Add(music domain.MusicToAdd) (int, error) {
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

func (r *MusicRepository) Delete(music domain.MusicToDelete) error {
	query := "DELETE FROM music WHERE song=$1"
	_, err := r.db.Exec(query, music.Song)
	return err
}

func (r *MusicRepository) Update(music domain.MusicToUpdate) error {
	query := "UPDATE music SET music_group=$1, text_song=$2, link=$3 WHERE song=$4"
	_, err := r.db.Exec(query, music.Group, music.Text, music.Link, music.Song)
	return err
}

func (r *MusicRepository) GetAll() ([]domain.Music, error) {
	var musics []domain.Music
	query := "SELECT music_group, song, text_song, link FROM music"
	if err := r.db.Select(&musics, query); err != nil {
		return nil, err
	}
	return musics, nil
}

func (r *MusicRepository) Get(music domain.MusicToGet) (domain.Music, error) {
	var foundMusic domain.Music
	query := "SELECT id, music_group, song, text_song, link FROM music WHERE song=$1"
	err := r.db.Get(&foundMusic, query, music.Song)
	return foundMusic, err
}
