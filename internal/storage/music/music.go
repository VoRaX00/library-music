package musicrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-music/internal/domain/models"
)

var (
	ErrMusicNotFound = errors.New("music not found")
)

type Music struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Music {
	return &Music{
		db: db,
	}
}

func (r *Music) Add(music models.Music) (int, error) {
	const op = "storage.music.Add"
	tx, err := r.db.Beginx()
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	err = r.insertGroup(tx, music.Group.Name)
	if err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	musicId, err := r.insertMusic(tx, music)
	if err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	query := `WITH group_cte AS (SELECT id FROM groups WHERE name = $1)
		INSERT INTO music_groups (music_id, group_id) 
		SELECT $2, g.id FROM group_cte g
		ON CONFLICT DO NOTHING;`

	_, err = tx.Exec(query, music.Group.Name, musicId)
	if err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return musicId, nil
}

func (r *Music) insertMusic(tx *sqlx.Tx, music models.Music) (int, error) {
	query := `INSERT INTO music (song, text_song, release_date, link) VALUES ($1, $2, $3, $4) RETURNING id;`

	var musicId int
	row := tx.QueryRow(query, music.Song, music.Text, music.ReleaseDate, music.Link)
	if err := row.Scan(&musicId); err != nil {
		return -1, err
	}
	return musicId, nil
}

func (r *Music) insertGroup(tx *sqlx.Tx, groupName string) error {
	query := `INSERT INTO groups (name) VALUES ($1) ON CONFLICT DO NOTHING;`
	_, err := tx.Exec(query, groupName)
	return err
}

func (r *Music) Delete(id int) error {
	const op = "storage.music.Delete"
	query := "DELETE FROM music WHERE id=$1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	return nil
}

func (r *Music) Update(music models.Music, id int) error {
	const op = "storage.music.Update"
	query := "UPDATE music SET song=$1, text_song=$2, release_date=$3, link=$4 WHERE id=$5"
	res, err := r.db.Exec(query, music.Song, music.Text, music.ReleaseDate, music.Link, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if row == 0 {
		return fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	return nil
}

func (r *Music) GetById(id int) (models.Music, error) {
	const op = "storage.music.GetById"
	var music models.Music
	query := `SELECT m.*, g.name
	FROM music m
	JOIN music_groups mg ON mg.music_id = m.id
	JOIN groups g ON g.id = mg.group_id
	WHERE m.id=?`

	err := r.db.Get(&music, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Music{}, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return models.Music{}, fmt.Errorf("%s: %w", op, err)
	}
	return music, nil
}

const pageSize = 5

func (r *Music) GetAll(params models.Music, page int) ([]models.Music, error) {
	const op = "storage.music.GetAll"
	var musics []models.Music
	query, args := generateQuery(params, page)
	err := r.db.Select(&musics, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(musics) == 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	return musics, nil
}

func generateQuery(params models.Music, page int) (string, []interface{}) {
	query := `SELECT id, song, link, release_date FROM music`
	var args []interface{}

	isWhere := false
	if params.Song != "" {
		args = append(args, params.Song)
		query += " WHERE song=" + fmt.Sprintf("$%d", len(args))
		isWhere = true
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
		args = append(args, params.Link)
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

func (r *Music) Get(song, group string) (models.Music, error) {
	const op = "storage.music.Get"

	var foundMusic models.Music
	query := `SELECT m.*, g.name
	FROM music m 
	JOIN music_groups mg ON m.id = mg.music_id 
    JOIN groups g ON mg.group_id = g.id 
	WHERE m.song = $1 AND g.name = $2`

	err := r.db.Get(&foundMusic, query, song, group)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Music{}, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return models.Music{}, fmt.Errorf("%s: %w", op, err)
	}
	return foundMusic, nil
}

func (r *Music) GetText(song, group string) (string, error) {
	const op = "storage.music.GetText"

	var text string
	query := `SELECT m.text_song
	FROM music m
	JOIN music_groups mg on m.id = mg.music_id
	JOIN groups g ON mg.group_id = g.id
	WHERE m.song = ? AND g.name = ?`
	err := r.db.Get(&text, query, song, group)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return text, err
}
