package musicrepo

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

func (r *Music) updateWithGroup(tx *sqlx.Tx, musicId int, group string) error {
	groupId, err := r.getGroupIDByName(tx, group)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if groupId == 0 {
		groupId, err = r.createGroup(tx, group)
		if err != nil {
			return err
		}
	}

	err = r.updateMusicGroup(tx, musicId, groupId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Music) getGroupIDByName(tx *sqlx.Tx, groupName string) (int, error) {
	query := `SELECT id FROM groups WHERE name = $1`
	var groupID int
	err := tx.Get(&groupID, query, groupName)
	if err != nil {
		return 0, err
	}
	return groupID, nil
}

func (r *Music) createGroup(tx *sqlx.Tx, groupName string) (int, error) {
	query := `INSERT INTO groups (name) VALUES ($1) RETURNING id;`
	var groupId int
	err := tx.QueryRow(query, groupName).Scan(&groupId)
	if err != nil {
		return 0, err
	}
	return groupId, nil
}

func (r *Music) updateMusicGroup(tx *sqlx.Tx, musicId, groupId int) error {
	query := `UPDATE music_groups SET group_id = $1 WHERE music_id = $2;`
	_, err := tx.Exec(query, groupId, musicId)
	if err != nil {
		return err
	}
	return nil
}
