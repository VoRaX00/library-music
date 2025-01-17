package musicrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-music/internal/domain/models"
	"reflect"
	"strings"
)

var (
	ErrMusicNotFound      = errors.New("music not found")
	ErrMusicAlreadyExists = errors.New("music already exists")
	ErrEmptyArguments     = errors.New("empty arguments")
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

	exists, err := r.checkSongInGroup(tx, music.Song, music.Group.Name)
	if err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	if exists {
		_ = tx.Rollback()
		return -1, fmt.Errorf("%s: %w", op, ErrMusicAlreadyExists)
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

func (r *Music) checkSongInGroup(tx *sqlx.Tx, song, groupName string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1
		FROM music m 
		JOIN music_groups mg ON m.id = mg.music_id
		JOIN groups g ON mg.group_id = g.id
		WHERE m.song = $1 AND g.name = $2
	)`

	var exists bool
	err := tx.Get(&exists, query, song, groupName)
	if err != nil {
		return false, err
	}
	return exists, nil
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
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if music.Group.Name != "" {
		err = r.updateWithGroup(tx, id, music.Group.Name)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	query, args := generateUpdateQuery(music, id)
	if args == nil {
		if music.Group.Name == "" {
			return fmt.Errorf("%s: %w", op, ErrEmptyArguments)
		}

		if err = tx.Commit(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}

	if music.Song != "" {
		res, err := r.checkUpdateOnDuplicate(tx, music.Song, id)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if res {
			return fmt.Errorf("%s: %w", op, ErrMusicAlreadyExists)
		}
	}

	res, err := tx.Exec(query, args...)
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

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func generateUpdateQuery(music models.Music, id int) (string, []interface{}) {
	query := "UPDATE music SET "
	var updates []string
	var args []interface{}
	v := reflect.ValueOf(music)
	t := reflect.TypeOf(music)

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() && t.Field(i).Name != "Group" {
			updates = append(updates, fmt.Sprintf("%s = $%d", t.Field(i).Tag.Get("db"), len(args)+1))
			args = append(args, v.Field(i).Interface())
		}
	}

	if len(updates) == 0 {
		return "", nil
	}

	query += strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id=$%d", len(args)+1)
	args = append(args, id)
	return query, args
}

func (r *Music) checkUpdateOnDuplicate(tx *sqlx.Tx, song string, id int) (bool, error) {
	query := `SELECT EXISTS (
    	SELECT 1
		FROM music m
		JOIN music_groups mg ON m.id = mg.music_id
		JOIN music_groups mg2 ON mg.group_id = mg2.group_id
		WHERE m.song = $1 AND mg2.music_id = $2 AND m.id <> $2
	)`

	var exists bool
	err := tx.Get(&exists, query, song, id)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *Music) GetById(id int) (models.Music, error) {
	const op = "storage.music.GetById"
	var music models.Music
	query := `SELECT m.id, m.song, m.text_song, m.link, m.release_date, g.name
	FROM music m
	JOIN music_groups mg ON mg.music_id = m.id
	JOIN groups g ON g.id = mg.group_id
	WHERE m.id=$1`

	err := r.db.Get(&music, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Music{}, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return models.Music{}, fmt.Errorf("%s: %w", op, err)
	}
	return music, nil
}

func (r *Music) GetAll(params models.Music, countSongs, page int) ([]models.Music, error) {
	const op = "storage.music.GetAll"
	var musics []models.Music
	query, args := generateQuery(params, countSongs, page)
	err := r.db.Select(&musics, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(musics) == 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	return musics, nil
}

func generateQuery(params models.Music, countSongs, page int) (string, []interface{}) {
	query := `SELECT m.id, m.song, m.text_song, m.link, m.release_date, 
       g.id AS "group.id",
       g.name AS "group.name"
       FROM music m
       LEFT JOIN music_groups mg ON mg.music_id = m.id
       LEFT JOIN groups g ON g.id = mg.group_id`

	var args []interface{}
	isWhere := false

	if params.Song != "" {
		args = append(args, params.Song)
		query += addCondition("m.song", len(args), false)
		isWhere = true
	}

	if params.Text != "" {
		args = append(args, params.Text)
		query += addCondition("m.text_song", len(args), isWhere)
		isWhere = true
	}

	if params.Link != "" {
		args = append(args, params.Link)
		query += addCondition("m.link", len(args), isWhere)
		isWhere = true
	}

	if params.Group.Name != "" {
		args = append(args, params.Group.Name)
		query += addCondition("g.name", len(args), isWhere)
		isWhere = true
	}

	if !params.ReleaseDate.IsZero() {
		args = append(args, params.ReleaseDate)
		query += addCondition("m.release_date", len(args), isWhere)
		isWhere = true
	}

	offset := (page - 1) * countSongs
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", countSongs, offset)
	return query, args
}

func addCondition(field string, paramIndex int, isWhere bool) string {
	condition := fmt.Sprintf("%s = $%d", field, paramIndex)
	if isWhere {
		return " AND " + condition
	}
	return " WHERE " + condition
}

func (r *Music) Get(song, group string) (models.Music, error) {
	const op = "storage.music.Get"

	var foundMusic models.Music
	query := `SELECT m.id, m.song, m.text_song, m.link, m.release_date, g.id AS "group.id",
        g.name AS "group.name"
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
	WHERE m.song = $1 AND g.name =$2`
	err := r.db.Get(&text, query, song, group)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return text, err
}
