// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// WsubSession represents a row from 'jlabo.wsub_session'.
type WsubSession struct {
	SessionID string         `json:"session_id"` // session_id
	ClientID  int8           `json:"client_id"`  // client_id
	Guest     sql.NullInt64  `json:"guest"`      // guest
	Time      sql.NullString `json:"time"`       // time
	Data      sql.NullString `json:"data"`       // data
	Userid    sql.NullInt64  `json:"userid"`     // userid
	Username  sql.NullString `json:"username"`   // username

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubSession exists in the database.
func (ws *WsubSession) Exists() bool {
	return ws._exists
}

// Deleted provides information if the WsubSession has been deleted from the database.
func (ws *WsubSession) Deleted() bool {
	return ws._deleted
}

// Insert inserts the WsubSession to the database.
func (ws *WsubSession) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ws._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO jlabo.wsub_session (` +
		`session_id, client_id, guest, time, data, userid, username` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ws.SessionID, ws.ClientID, ws.Guest, ws.Time, ws.Data, ws.Userid, ws.Username)
	_, err = db.Exec(sqlstr, ws.SessionID, ws.ClientID, ws.Guest, ws.Time, ws.Data, ws.Userid, ws.Username)
	if err != nil {
		return err
	}

	// set existence
	ws._exists = true

	return nil
}

// Update updates the WsubSession in the database.
func (ws *WsubSession) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ws._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ws._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_session SET ` +
		`client_id = ?, guest = ?, time = ?, data = ?, userid = ?, username = ?` +
		` WHERE session_id = ?`

	// run query
	XOLog(sqlstr, ws.ClientID, ws.Guest, ws.Time, ws.Data, ws.Userid, ws.Username, ws.SessionID)
	_, err = db.Exec(sqlstr, ws.ClientID, ws.Guest, ws.Time, ws.Data, ws.Userid, ws.Username, ws.SessionID)
	return err
}

// Save saves the WsubSession to the database.
func (ws *WsubSession) Save(db XODB) error {
	if ws.Exists() {
		return ws.Update(db)
	}

	return ws.Insert(db)
}

// Delete deletes the WsubSession from the database.
func (ws *WsubSession) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ws._exists {
		return nil
	}

	// if deleted, bail
	if ws._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_session WHERE session_id = ?`

	// run query
	XOLog(sqlstr, ws.SessionID)
	_, err = db.Exec(sqlstr, ws.SessionID)
	if err != nil {
		return err
	}

	// set deleted
	ws._deleted = true

	return nil
}

// WsubSessionsByTime retrieves a row from 'jlabo.wsub_session' as a WsubSession.
//
// Generated from index 'time'.
func WsubSessionsByTime(db XODB, time sql.NullString) ([]*WsubSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_id, client_id, guest, time, data, userid, username ` +
		`FROM jlabo.wsub_session ` +
		`WHERE time = ?`

	// run query
	XOLog(sqlstr, time)
	q, err := db.Query(sqlstr, time)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubSession{}
	for q.Next() {
		ws := WsubSession{
			_exists: true,
		}

		// scan
		err = q.Scan(&ws.SessionID, &ws.ClientID, &ws.Guest, &ws.Time, &ws.Data, &ws.Userid, &ws.Username)
		if err != nil {
			return nil, err
		}

		res = append(res, &ws)
	}

	return res, nil
}

// WsubSessionsByUserid retrieves a row from 'jlabo.wsub_session' as a WsubSession.
//
// Generated from index 'userid'.
func WsubSessionsByUserid(db XODB, userid sql.NullInt64) ([]*WsubSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_id, client_id, guest, time, data, userid, username ` +
		`FROM jlabo.wsub_session ` +
		`WHERE userid = ?`

	// run query
	XOLog(sqlstr, userid)
	q, err := db.Query(sqlstr, userid)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubSession{}
	for q.Next() {
		ws := WsubSession{
			_exists: true,
		}

		// scan
		err = q.Scan(&ws.SessionID, &ws.ClientID, &ws.Guest, &ws.Time, &ws.Data, &ws.Userid, &ws.Username)
		if err != nil {
			return nil, err
		}

		res = append(res, &ws)
	}

	return res, nil
}

// WsubSessionBySessionID retrieves a row from 'jlabo.wsub_session' as a WsubSession.
//
// Generated from index 'wsub_session_session_id_pkey'.
func WsubSessionBySessionID(db XODB, sessionID string) (*WsubSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_id, client_id, guest, time, data, userid, username ` +
		`FROM jlabo.wsub_session ` +
		`WHERE session_id = ?`

	// run query
	XOLog(sqlstr, sessionID)
	ws := WsubSession{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, sessionID).Scan(&ws.SessionID, &ws.ClientID, &ws.Guest, &ws.Time, &ws.Data, &ws.Userid, &ws.Username)
	if err != nil {
		return nil, err
	}

	return &ws, nil
}
