package infrastructure

import (
	"fmt"
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
	query := "INSERT INTO music (music_group, song, text_song, link, release_date) values ($1, $2, $3, $4, $5) RETURNING id"

	row := tx.QueryRow(query, music.Group, music.Song, music.Text, music.Link, music.ReleaseDate)
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

func (r *MusicRepository) Update(music domain.Music, id int) error {
	query := "UPDATE music SET song=$1, music_group=$2, text_song=$3, link=$4, release_date=$5 WHERE id=$6"
	_, err := r.db.Exec(query, music.Song, music.Group, music.Text, music.Link, music.ReleaseDate, id)
	return err
}

func generateQuery(params domain.Music, page int) (string, []interface{}) {
	query := "SELECT id, music_group, song, link, to_char(release_date, 'DD-MM-YYYY') as release_date FROM music"
	var args []interface{}

	isWhere := false
	if params.Song != "" {
		args = append(args, params.Song)
		query += " WHERE song=" + fmt.Sprintf("$%d", len(args))
		isWhere = true
	}

	if params.Group != "" {
		args = append(args, params.Group)
		if isWhere {
			query += " AND music_group=" + fmt.Sprintf("$%d", len(args))
		} else {
			query += " WHERE music_group=" + fmt.Sprintf("$%d", len(args))
			isWhere = true
		}
	}

	if params.Text != "" {
		args = append(args, params.Text)
		if isWhere {
			query += " AND text_song=" + fmt.Sprintf("$%d", len(args))
		} else {
			query += " WHERE text_song=" + fmt.Sprintf("$%d", len(args))
			isWhere = true
		}
	}

	if params.Link != "" {
		args = append(args, params.Text)
		if isWhere {
			query += " AND link=" + fmt.Sprintf("$%d", len(args))
		} else {
			query += " WHERE link=" + fmt.Sprintf("$%d", len(args))
			isWhere = true
		}
	}

	offset := (page - 1) * pageSize

	query += " LIMIT " + fmt.Sprintf("%d", pageSize) + " OFFSET " + fmt.Sprintf("%d", offset)
	return query, args
}

const pageSize = 5

func (r *MusicRepository) GetAll(params domain.Music, page int) ([]domain.Music, error) {
	var musics []domain.Music
	query, args := generateQuery(params, page)
	if err := r.db.Select(&musics, query, args...); err != nil {
		return nil, err
	}
	return musics, nil
}

func (r *MusicRepository) Get(song, group string) (domain.Music, error) {
	var foundMusic domain.Music
	query := `SELECT id, music_group, song, link, to_char(release_date, 'DD-MM-YYYY') as release_date
		FROM music WHERE song=$1 AND music_group=$2`
	err := r.db.Get(&foundMusic, query, song, group)
	return foundMusic, err
}

func (r *MusicRepository) GetText(song, group string) (string, error) {
	var text string
	query := "SELECT text_song FROM music WHERE song=$1 AND music_group=$2"
	err := r.db.Get(&text, query, song, group)
	return text, err
}
