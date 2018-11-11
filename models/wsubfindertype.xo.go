// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubFinderType represents a row from 'jlabo.wsub_finder_types'.
type WsubFinderType struct {
	ID    uint   `json:"id"`    // id
	Title string `json:"title"` // title
	Mime  string `json:"mime"`  // mime

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderType exists in the database.
func (wft *WsubFinderType) Exists() bool {
	return wft._exists
}

// Deleted provides information if the WsubFinderType has been deleted from the database.
func (wft *WsubFinderType) Deleted() bool {
	return wft._deleted
}

// Insert inserts the WsubFinderType to the database.
func (wft *WsubFinderType) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wft._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_finder_types (` +
		`title, mime` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wft.Title, wft.Mime)
	res, err := db.Exec(sqlstr, wft.Title, wft.Mime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wft.ID = uint(id)
	wft._exists = true

	return nil
}

// Update updates the WsubFinderType in the database.
func (wft *WsubFinderType) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wft._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wft._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_finder_types SET ` +
		`title = ?, mime = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wft.Title, wft.Mime, wft.ID)
	_, err = db.Exec(sqlstr, wft.Title, wft.Mime, wft.ID)
	return err
}

// Save saves the WsubFinderType to the database.
func (wft *WsubFinderType) Save(db XODB) error {
	if wft.Exists() {
		return wft.Update(db)
	}

	return wft.Insert(db)
}

// Delete deletes the WsubFinderType from the database.
func (wft *WsubFinderType) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wft._exists {
		return nil
	}

	// if deleted, bail
	if wft._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_finder_types WHERE id = ?`

	// run query
	XOLog(sqlstr, wft.ID)
	_, err = db.Exec(sqlstr, wft.ID)
	if err != nil {
		return err
	}

	// set deleted
	wft._deleted = true

	return nil
}

// WsubFinderTypeByTitle retrieves a row from 'jlabo.wsub_finder_types' as a WsubFinderType.
//
// Generated from index 'title'.
func WsubFinderTypeByTitle(db XODB, title string) (*WsubFinderType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, mime ` +
		`FROM jlabo.wsub_finder_types ` +
		`WHERE title = ?`

	// run query
	XOLog(sqlstr, title)
	wft := WsubFinderType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, title).Scan(&wft.ID, &wft.Title, &wft.Mime)
	if err != nil {
		return nil, err
	}

	return &wft, nil
}

// WsubFinderTypeByID retrieves a row from 'jlabo.wsub_finder_types' as a WsubFinderType.
//
// Generated from index 'wsub_finder_types_id_pkey'.
func WsubFinderTypeByID(db XODB, id uint) (*WsubFinderType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, mime ` +
		`FROM jlabo.wsub_finder_types ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wft := WsubFinderType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wft.ID, &wft.Title, &wft.Mime)
	if err != nil {
		return nil, err
	}

	return &wft, nil
}