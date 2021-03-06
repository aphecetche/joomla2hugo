// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// WsubUcmHistory represents a row from 'jlabo.wsub_ucm_history'.
type WsubUcmHistory struct {
	VersionID      uint      `json:"version_id"`      // version_id
	UcmItemID      uint      `json:"ucm_item_id"`     // ucm_item_id
	UcmTypeID      uint      `json:"ucm_type_id"`     // ucm_type_id
	VersionNote    string    `json:"version_note"`    // version_note
	SaveDate       time.Time `json:"save_date"`       // save_date
	EditorUserID   uint      `json:"editor_user_id"`  // editor_user_id
	CharacterCount uint      `json:"character_count"` // character_count
	Sha1Hash       string    `json:"sha1_hash"`       // sha1_hash
	VersionData    string    `json:"version_data"`    // version_data
	KeepForever    int8      `json:"keep_forever"`    // keep_forever

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubUcmHistory exists in the database.
func (wuh *WsubUcmHistory) Exists() bool {
	return wuh._exists
}

// Deleted provides information if the WsubUcmHistory has been deleted from the database.
func (wuh *WsubUcmHistory) Deleted() bool {
	return wuh._deleted
}

// Insert inserts the WsubUcmHistory to the database.
func (wuh *WsubUcmHistory) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wuh._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_ucm_history (` +
		`ucm_item_id, ucm_type_id, version_note, save_date, editor_user_id, character_count, sha1_hash, version_data, keep_forever` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wuh.UcmItemID, wuh.UcmTypeID, wuh.VersionNote, wuh.SaveDate, wuh.EditorUserID, wuh.CharacterCount, wuh.Sha1Hash, wuh.VersionData, wuh.KeepForever)
	res, err := db.Exec(sqlstr, wuh.UcmItemID, wuh.UcmTypeID, wuh.VersionNote, wuh.SaveDate, wuh.EditorUserID, wuh.CharacterCount, wuh.Sha1Hash, wuh.VersionData, wuh.KeepForever)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wuh.VersionID = uint(id)
	wuh._exists = true

	return nil
}

// Update updates the WsubUcmHistory in the database.
func (wuh *WsubUcmHistory) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wuh._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wuh._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_ucm_history SET ` +
		`ucm_item_id = ?, ucm_type_id = ?, version_note = ?, save_date = ?, editor_user_id = ?, character_count = ?, sha1_hash = ?, version_data = ?, keep_forever = ?` +
		` WHERE version_id = ?`

	// run query
	XOLog(sqlstr, wuh.UcmItemID, wuh.UcmTypeID, wuh.VersionNote, wuh.SaveDate, wuh.EditorUserID, wuh.CharacterCount, wuh.Sha1Hash, wuh.VersionData, wuh.KeepForever, wuh.VersionID)
	_, err = db.Exec(sqlstr, wuh.UcmItemID, wuh.UcmTypeID, wuh.VersionNote, wuh.SaveDate, wuh.EditorUserID, wuh.CharacterCount, wuh.Sha1Hash, wuh.VersionData, wuh.KeepForever, wuh.VersionID)
	return err
}

// Save saves the WsubUcmHistory to the database.
func (wuh *WsubUcmHistory) Save(db XODB) error {
	if wuh.Exists() {
		return wuh.Update(db)
	}

	return wuh.Insert(db)
}

// Delete deletes the WsubUcmHistory from the database.
func (wuh *WsubUcmHistory) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wuh._exists {
		return nil
	}

	// if deleted, bail
	if wuh._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_ucm_history WHERE version_id = ?`

	// run query
	XOLog(sqlstr, wuh.VersionID)
	_, err = db.Exec(sqlstr, wuh.VersionID)
	if err != nil {
		return err
	}

	// set deleted
	wuh._deleted = true

	return nil
}

// WsubUcmHistoriesBySaveDate retrieves a row from 'jlabo.wsub_ucm_history' as a WsubUcmHistory.
//
// Generated from index 'idx_save_date'.
func WsubUcmHistoriesBySaveDate(db XODB, saveDate time.Time) ([]*WsubUcmHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`version_id, ucm_item_id, ucm_type_id, version_note, save_date, editor_user_id, character_count, sha1_hash, version_data, keep_forever ` +
		`FROM jlabo.wsub_ucm_history ` +
		`WHERE save_date = ?`

	// run query
	XOLog(sqlstr, saveDate)
	q, err := db.Query(sqlstr, saveDate)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubUcmHistory{}
	for q.Next() {
		wuh := WsubUcmHistory{
			_exists: true,
		}

		// scan
		err = q.Scan(&wuh.VersionID, &wuh.UcmItemID, &wuh.UcmTypeID, &wuh.VersionNote, &wuh.SaveDate, &wuh.EditorUserID, &wuh.CharacterCount, &wuh.Sha1Hash, &wuh.VersionData, &wuh.KeepForever)
		if err != nil {
			return nil, err
		}

		res = append(res, &wuh)
	}

	return res, nil
}

// WsubUcmHistoriesByUcmTypeIDUcmItemID retrieves a row from 'jlabo.wsub_ucm_history' as a WsubUcmHistory.
//
// Generated from index 'idx_ucm_item_id'.
func WsubUcmHistoriesByUcmTypeIDUcmItemID(db XODB, ucmTypeID uint, ucmItemID uint) ([]*WsubUcmHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`version_id, ucm_item_id, ucm_type_id, version_note, save_date, editor_user_id, character_count, sha1_hash, version_data, keep_forever ` +
		`FROM jlabo.wsub_ucm_history ` +
		`WHERE ucm_type_id = ? AND ucm_item_id = ?`

	// run query
	XOLog(sqlstr, ucmTypeID, ucmItemID)
	q, err := db.Query(sqlstr, ucmTypeID, ucmItemID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubUcmHistory{}
	for q.Next() {
		wuh := WsubUcmHistory{
			_exists: true,
		}

		// scan
		err = q.Scan(&wuh.VersionID, &wuh.UcmItemID, &wuh.UcmTypeID, &wuh.VersionNote, &wuh.SaveDate, &wuh.EditorUserID, &wuh.CharacterCount, &wuh.Sha1Hash, &wuh.VersionData, &wuh.KeepForever)
		if err != nil {
			return nil, err
		}

		res = append(res, &wuh)
	}

	return res, nil
}

// WsubUcmHistoryByVersionID retrieves a row from 'jlabo.wsub_ucm_history' as a WsubUcmHistory.
//
// Generated from index 'wsub_ucm_history_version_id_pkey'.
func WsubUcmHistoryByVersionID(db XODB, versionID uint) (*WsubUcmHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`version_id, ucm_item_id, ucm_type_id, version_note, save_date, editor_user_id, character_count, sha1_hash, version_data, keep_forever ` +
		`FROM jlabo.wsub_ucm_history ` +
		`WHERE version_id = ?`

	// run query
	XOLog(sqlstr, versionID)
	wuh := WsubUcmHistory{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, versionID).Scan(&wuh.VersionID, &wuh.UcmItemID, &wuh.UcmTypeID, &wuh.VersionNote, &wuh.SaveDate, &wuh.EditorUserID, &wuh.CharacterCount, &wuh.Sha1Hash, &wuh.VersionData, &wuh.KeepForever)
	if err != nil {
		return nil, err
	}

	return &wuh, nil
}
