// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// WsubFcWorkflow represents a row from 'jlabo.wsub_fc_workflow'.
type WsubFcWorkflow struct {
	ID        int            `json:"id"`         // id
	Notes     sql.NullString `json:"notes"`      // notes
	VersionID sql.NullInt64  `json:"version_id"` // version_id
	Stage     sql.NullInt64  `json:"stage"`      // stage
	ContentID sql.NullInt64  `json:"content_id"` // content_id
	Reviewer  sql.NullString `json:"reviewer"`   // reviewer
	Ts        mysql.NullTime `json:"ts"`         // ts

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFcWorkflow exists in the database.
func (wfw *WsubFcWorkflow) Exists() bool {
	return wfw._exists
}

// Deleted provides information if the WsubFcWorkflow has been deleted from the database.
func (wfw *WsubFcWorkflow) Deleted() bool {
	return wfw._deleted
}

// Insert inserts the WsubFcWorkflow to the database.
func (wfw *WsubFcWorkflow) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wfw._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_fc_workflow (` +
		`notes, version_id, stage, content_id, reviewer, ts` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wfw.Notes, wfw.VersionID, wfw.Stage, wfw.ContentID, wfw.Reviewer, wfw.Ts)
	res, err := db.Exec(sqlstr, wfw.Notes, wfw.VersionID, wfw.Stage, wfw.ContentID, wfw.Reviewer, wfw.Ts)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wfw.ID = int(id)
	wfw._exists = true

	return nil
}

// Update updates the WsubFcWorkflow in the database.
func (wfw *WsubFcWorkflow) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wfw._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wfw._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_fc_workflow SET ` +
		`notes = ?, version_id = ?, stage = ?, content_id = ?, reviewer = ?, ts = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wfw.Notes, wfw.VersionID, wfw.Stage, wfw.ContentID, wfw.Reviewer, wfw.Ts, wfw.ID)
	_, err = db.Exec(sqlstr, wfw.Notes, wfw.VersionID, wfw.Stage, wfw.ContentID, wfw.Reviewer, wfw.Ts, wfw.ID)
	return err
}

// Save saves the WsubFcWorkflow to the database.
func (wfw *WsubFcWorkflow) Save(db XODB) error {
	if wfw.Exists() {
		return wfw.Update(db)
	}

	return wfw.Insert(db)
}

// Delete deletes the WsubFcWorkflow from the database.
func (wfw *WsubFcWorkflow) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wfw._exists {
		return nil
	}

	// if deleted, bail
	if wfw._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_fc_workflow WHERE id = ?`

	// run query
	XOLog(sqlstr, wfw.ID)
	_, err = db.Exec(sqlstr, wfw.ID)
	if err != nil {
		return err
	}

	// set deleted
	wfw._deleted = true

	return nil
}

// WsubFcWorkflowByID retrieves a row from 'jlabo.wsub_fc_workflow' as a WsubFcWorkflow.
//
// Generated from index 'wsub_fc_workflow_id_pkey'.
func WsubFcWorkflowByID(db XODB, id int) (*WsubFcWorkflow, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, notes, version_id, stage, content_id, reviewer, ts ` +
		`FROM jlabo.wsub_fc_workflow ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wfw := WsubFcWorkflow{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wfw.ID, &wfw.Notes, &wfw.VersionID, &wfw.Stage, &wfw.ContentID, &wfw.Reviewer, &wfw.Ts)
	if err != nil {
		return nil, err
	}

	return &wfw, nil
}
