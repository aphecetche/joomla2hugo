// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubFinderTerm represents a row from 'jlabo.wsub_finder_terms'.
type WsubFinderTerm struct {
	TermID   uint    `json:"term_id"`  // term_id
	Term     string  `json:"term"`     // term
	Stem     string  `json:"stem"`     // stem
	Common   bool    `json:"common"`   // common
	Phrase   bool    `json:"phrase"`   // phrase
	Weight   float32 `json:"weight"`   // weight
	Soundex  string  `json:"soundex"`  // soundex
	Links    int     `json:"links"`    // links
	Language string  `json:"language"` // language

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderTerm exists in the database.
func (wft *WsubFinderTerm) Exists() bool {
	return wft._exists
}

// Deleted provides information if the WsubFinderTerm has been deleted from the database.
func (wft *WsubFinderTerm) Deleted() bool {
	return wft._deleted
}

// Insert inserts the WsubFinderTerm to the database.
func (wft *WsubFinderTerm) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wft._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_finder_terms (` +
		`term, stem, common, phrase, weight, soundex, links, language` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wft.Term, wft.Stem, wft.Common, wft.Phrase, wft.Weight, wft.Soundex, wft.Links, wft.Language)
	res, err := db.Exec(sqlstr, wft.Term, wft.Stem, wft.Common, wft.Phrase, wft.Weight, wft.Soundex, wft.Links, wft.Language)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wft.TermID = uint(id)
	wft._exists = true

	return nil
}

// Update updates the WsubFinderTerm in the database.
func (wft *WsubFinderTerm) Update(db XODB) error {
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
	const sqlstr = `UPDATE jlabo.wsub_finder_terms SET ` +
		`term = ?, stem = ?, common = ?, phrase = ?, weight = ?, soundex = ?, links = ?, language = ?` +
		` WHERE term_id = ?`

	// run query
	XOLog(sqlstr, wft.Term, wft.Stem, wft.Common, wft.Phrase, wft.Weight, wft.Soundex, wft.Links, wft.Language, wft.TermID)
	_, err = db.Exec(sqlstr, wft.Term, wft.Stem, wft.Common, wft.Phrase, wft.Weight, wft.Soundex, wft.Links, wft.Language, wft.TermID)
	return err
}

// Save saves the WsubFinderTerm to the database.
func (wft *WsubFinderTerm) Save(db XODB) error {
	if wft.Exists() {
		return wft.Update(db)
	}

	return wft.Insert(db)
}

// Delete deletes the WsubFinderTerm from the database.
func (wft *WsubFinderTerm) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM jlabo.wsub_finder_terms WHERE term_id = ?`

	// run query
	XOLog(sqlstr, wft.TermID)
	_, err = db.Exec(sqlstr, wft.TermID)
	if err != nil {
		return err
	}

	// set deleted
	wft._deleted = true

	return nil
}

// WsubFinderTermsBySoundexPhrase retrieves a row from 'jlabo.wsub_finder_terms' as a WsubFinderTerm.
//
// Generated from index 'idx_soundex_phrase'.
func WsubFinderTermsBySoundexPhrase(db XODB, soundex string, phrase bool) ([]*WsubFinderTerm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term_id, term, stem, common, phrase, weight, soundex, links, language ` +
		`FROM jlabo.wsub_finder_terms ` +
		`WHERE soundex = ? AND phrase = ?`

	// run query
	XOLog(sqlstr, soundex, phrase)
	q, err := db.Query(sqlstr, soundex, phrase)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTerm{}
	for q.Next() {
		wft := WsubFinderTerm{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.TermID, &wft.Term, &wft.Stem, &wft.Common, &wft.Phrase, &wft.Weight, &wft.Soundex, &wft.Links, &wft.Language)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTermsByStemPhrase retrieves a row from 'jlabo.wsub_finder_terms' as a WsubFinderTerm.
//
// Generated from index 'idx_stem_phrase'.
func WsubFinderTermsByStemPhrase(db XODB, stem string, phrase bool) ([]*WsubFinderTerm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term_id, term, stem, common, phrase, weight, soundex, links, language ` +
		`FROM jlabo.wsub_finder_terms ` +
		`WHERE stem = ? AND phrase = ?`

	// run query
	XOLog(sqlstr, stem, phrase)
	q, err := db.Query(sqlstr, stem, phrase)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTerm{}
	for q.Next() {
		wft := WsubFinderTerm{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.TermID, &wft.Term, &wft.Stem, &wft.Common, &wft.Phrase, &wft.Weight, &wft.Soundex, &wft.Links, &wft.Language)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTermByTerm retrieves a row from 'jlabo.wsub_finder_terms' as a WsubFinderTerm.
//
// Generated from index 'idx_term'.
func WsubFinderTermByTerm(db XODB, term string) (*WsubFinderTerm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term_id, term, stem, common, phrase, weight, soundex, links, language ` +
		`FROM jlabo.wsub_finder_terms ` +
		`WHERE term = ?`

	// run query
	XOLog(sqlstr, term)
	wft := WsubFinderTerm{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, term).Scan(&wft.TermID, &wft.Term, &wft.Stem, &wft.Common, &wft.Phrase, &wft.Weight, &wft.Soundex, &wft.Links, &wft.Language)
	if err != nil {
		return nil, err
	}

	return &wft, nil
}

// WsubFinderTermsByTermPhrase retrieves a row from 'jlabo.wsub_finder_terms' as a WsubFinderTerm.
//
// Generated from index 'idx_term_phrase'.
func WsubFinderTermsByTermPhrase(db XODB, term string, phrase bool) ([]*WsubFinderTerm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term_id, term, stem, common, phrase, weight, soundex, links, language ` +
		`FROM jlabo.wsub_finder_terms ` +
		`WHERE term = ? AND phrase = ?`

	// run query
	XOLog(sqlstr, term, phrase)
	q, err := db.Query(sqlstr, term, phrase)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTerm{}
	for q.Next() {
		wft := WsubFinderTerm{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.TermID, &wft.Term, &wft.Stem, &wft.Common, &wft.Phrase, &wft.Weight, &wft.Soundex, &wft.Links, &wft.Language)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTermByTermID retrieves a row from 'jlabo.wsub_finder_terms' as a WsubFinderTerm.
//
// Generated from index 'wsub_finder_terms_term_id_pkey'.
func WsubFinderTermByTermID(db XODB, termID uint) (*WsubFinderTerm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term_id, term, stem, common, phrase, weight, soundex, links, language ` +
		`FROM jlabo.wsub_finder_terms ` +
		`WHERE term_id = ?`

	// run query
	XOLog(sqlstr, termID)
	wft := WsubFinderTerm{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, termID).Scan(&wft.TermID, &wft.Term, &wft.Stem, &wft.Common, &wft.Phrase, &wft.Weight, &wft.Soundex, &wft.Links, &wft.Language)
	if err != nil {
		return nil, err
	}

	return &wft, nil
}
