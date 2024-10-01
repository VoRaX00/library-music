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

func (r *MusicRepository) Delete(id int) error {
	query := "DELETE FROM music WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MusicRepository) Update(music domain.MusicToUpdate, id int) error {
	query := "UPDATE music SET song=$1, music_group=$2, text_song=$3, link=$4 WHERE id=$5"
	_, err := r.db.Exec(query, music.Text, music.Link, music.Song, music.Group, id)
	return err
}

const pageSize = 5

func (r *MusicRepository) GetAll(page int) ([]domain.MusicToGet, error) {
	var musics []domain.MusicToGet
	query := "SELECT id, music_group, song, link FROM music LIMIT $1 OFFSET $2"
	offset := (page - 1) * pageSize
	if err := r.db.Select(&musics, query, pageSize, offset); err != nil {
		return nil, err
	}
	return musics, nil
}

func (r *MusicRepository) Get(song, group string) (domain.MusicToGet, error) {
	var foundMusic domain.MusicToGet
	query := "SELECT id, music_group, song, link FROM music WHERE song=$1 AND music_group=$2"
	err := r.db.Get(&foundMusic, query, song, group)
	return foundMusic, err
}

func (r *MusicRepository) GetText(song, group string) (string, error) {
	var text string
	query := "SELECT text_song FROM music WHERE song=$1 AND music_group=$2"
	err := r.db.Get(&text, query, song, group)
	return text, err
}
