package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"library-music/internal/domain/models"
)

type Music struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Music {
	return &Music{
		db: db,
	}
}

func (r *Music) getGroupId(name string) (int, error) {
	var id int
	query := "SELECT id FROM groups WHERE name=?"
	row := r.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Music) insertMusic(tx *sqlx.Tx, music models.Music) (int, error) {
	var musicId int
	query := `INSERT INTO music (song, text_song, release_date, link) VALUES (?, ?, ?, ?);`
	row := tx.QueryRow(query, music.Song, music.Text, music.ReleaseDate, music.Link)

	if err := row.Scan(&musicId); err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	return musicId, nil
}

func (r *Music) insertGroup(tx *sqlx.Tx, groupName string) (int, error) {
	groupId, err := r.getGroupId(groupName)
	if err == nil {
		return groupId, nil
	}

	query := `INSERT INTO groups (name) VALUES (?)`
	row := tx.QueryRow(query, groupName)
	if err = row.Scan(&groupId); err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	return groupId, nil
}

func (r *Music) insertMusicGroups(tx *sqlx.Tx, musicId, groupId int) error {
	query := `INSERT INTO music_groups (music_id, group_id) VALUES (?, ?)`
	row := tx.QueryRow(query, musicId, groupId)
	if row.Err() != nil {
		_ = tx.Rollback()
		return row.Err()
	}
	return nil
}

func (r *Music) Add(music models.Music) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return -1, err
	}

	groupId, err := r.insertGroup(tx, music.Group)
	musicId, err := r.insertMusic(tx, music)

	err = r.insertMusicGroups(tx, musicId, groupId)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	return musicId, tx.Commit()
}

func (r *Music) Delete(id int) error {
	query := "DELETE FROM music_groups WHERE music_id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM music WHERE id=$1"
	_, err = r.db.Exec(query, id)
	return err
}

func (r *Music) Update(music models.Music, id int) (models.Music, error) {
	query := "UPDATE music SET song=?, text_song=?, release_date=?, link=? WHERE id=?"
	_, err := r.db.Exec(query, music.Song, music.Text, music.ReleaseDate, music.Link, id)
	return models.Music{}, err
}

func generateQuery(params models.Music, page int) (string, []interface{}) {
	query := "SELECT id, song, link, release_date FROM music"
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

func (r *Music) GetById(id int) (models.Music, error) {
	var music models.Music
	query := `SELECT m.*, g.name
	FROM music m
	JOIN music_groups mg ON mg.music_id = m.id
	JOIN groups g ON g.id = mg.group_id
	WHERE m.id=?`

	err := r.db.Select(&music, query, id)
	return music, err
}

const pageSize = 5

func (r *Music) GetAll(params models.Music, page int) ([]models.Music, error) {
	var musics []models.Music
	query, args := generateQuery(params, page)
	err := r.db.Select(&musics, query, args...)
	return musics, err
}

func (r *Music) Get(song, group string) (models.Music, error) {
	var foundMusic models.Music
	query := `SELECT m.*, g.name
	FROM music m 
	JOIN music_groups mg ON m.id = mg.music_id 
    JOIN groups g ON mg.group_id = g.id 
	WHERE m.song = ? AND g.name = ?`

	err := r.db.Get(&foundMusic, query, song, group)
	return foundMusic, err
}

func (r *Music) GetText(song, group string) (string, error) {
	var text string
	query := `SELECT m.text_song
	FROM music m
	JOIN music_groups mg on m.id = mg.music_id
	JOIN groups g ON mg.group_id = g.id
	WHERE m.song = ? AND g.name = ?`
	err := r.db.Get(&text, query, song, group)
	return text, err
}
