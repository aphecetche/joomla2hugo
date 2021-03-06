// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// WsubPostinstallMessage represents a row from 'jlabo.wsub_postinstall_messages'.
type WsubPostinstallMessage struct {
	PostinstallMessageID uint64         `json:"postinstall_message_id"` // postinstall_message_id
	ExtensionID          int64          `json:"extension_id"`           // extension_id
	TitleKey             string         `json:"title_key"`              // title_key
	DescriptionKey       string         `json:"description_key"`        // description_key
	ActionKey            string         `json:"action_key"`             // action_key
	LanguageExtension    string         `json:"language_extension"`     // language_extension
	LanguageClientID     int8           `json:"language_client_id"`     // language_client_id
	Type                 string         `json:"type"`                   // type
	ActionFile           sql.NullString `json:"action_file"`            // action_file
	Action               sql.NullString `json:"action"`                 // action
	ConditionFile        sql.NullString `json:"condition_file"`         // condition_file
	ConditionMethod      sql.NullString `json:"condition_method"`       // condition_method
	VersionIntroduced    string         `json:"version_introduced"`     // version_introduced
	Enabled              int8           `json:"enabled"`                // enabled

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubPostinstallMessage exists in the database.
func (wpm *WsubPostinstallMessage) Exists() bool {
	return wpm._exists
}

// Deleted provides information if the WsubPostinstallMessage has been deleted from the database.
func (wpm *WsubPostinstallMessage) Deleted() bool {
	return wpm._deleted
}

// Insert inserts the WsubPostinstallMessage to the database.
func (wpm *WsubPostinstallMessage) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wpm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_postinstall_messages (` +
		`extension_id, title_key, description_key, action_key, language_extension, language_client_id, type, action_file, action, condition_file, condition_method, version_introduced, enabled` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wpm.ExtensionID, wpm.TitleKey, wpm.DescriptionKey, wpm.ActionKey, wpm.LanguageExtension, wpm.LanguageClientID, wpm.Type, wpm.ActionFile, wpm.Action, wpm.ConditionFile, wpm.ConditionMethod, wpm.VersionIntroduced, wpm.Enabled)
	res, err := db.Exec(sqlstr, wpm.ExtensionID, wpm.TitleKey, wpm.DescriptionKey, wpm.ActionKey, wpm.LanguageExtension, wpm.LanguageClientID, wpm.Type, wpm.ActionFile, wpm.Action, wpm.ConditionFile, wpm.ConditionMethod, wpm.VersionIntroduced, wpm.Enabled)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wpm.PostinstallMessageID = uint64(id)
	wpm._exists = true

	return nil
}

// Update updates the WsubPostinstallMessage in the database.
func (wpm *WsubPostinstallMessage) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wpm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wpm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_postinstall_messages SET ` +
		`extension_id = ?, title_key = ?, description_key = ?, action_key = ?, language_extension = ?, language_client_id = ?, type = ?, action_file = ?, action = ?, condition_file = ?, condition_method = ?, version_introduced = ?, enabled = ?` +
		` WHERE postinstall_message_id = ?`

	// run query
	XOLog(sqlstr, wpm.ExtensionID, wpm.TitleKey, wpm.DescriptionKey, wpm.ActionKey, wpm.LanguageExtension, wpm.LanguageClientID, wpm.Type, wpm.ActionFile, wpm.Action, wpm.ConditionFile, wpm.ConditionMethod, wpm.VersionIntroduced, wpm.Enabled, wpm.PostinstallMessageID)
	_, err = db.Exec(sqlstr, wpm.ExtensionID, wpm.TitleKey, wpm.DescriptionKey, wpm.ActionKey, wpm.LanguageExtension, wpm.LanguageClientID, wpm.Type, wpm.ActionFile, wpm.Action, wpm.ConditionFile, wpm.ConditionMethod, wpm.VersionIntroduced, wpm.Enabled, wpm.PostinstallMessageID)
	return err
}

// Save saves the WsubPostinstallMessage to the database.
func (wpm *WsubPostinstallMessage) Save(db XODB) error {
	if wpm.Exists() {
		return wpm.Update(db)
	}

	return wpm.Insert(db)
}

// Delete deletes the WsubPostinstallMessage from the database.
func (wpm *WsubPostinstallMessage) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wpm._exists {
		return nil
	}

	// if deleted, bail
	if wpm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_postinstall_messages WHERE postinstall_message_id = ?`

	// run query
	XOLog(sqlstr, wpm.PostinstallMessageID)
	_, err = db.Exec(sqlstr, wpm.PostinstallMessageID)
	if err != nil {
		return err
	}

	// set deleted
	wpm._deleted = true

	return nil
}

// WsubPostinstallMessageByPostinstallMessageID retrieves a row from 'jlabo.wsub_postinstall_messages' as a WsubPostinstallMessage.
//
// Generated from index 'wsub_postinstall_messages_postinstall_message_id_pkey'.
func WsubPostinstallMessageByPostinstallMessageID(db XODB, postinstallMessageID uint64) (*WsubPostinstallMessage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`postinstall_message_id, extension_id, title_key, description_key, action_key, language_extension, language_client_id, type, action_file, action, condition_file, condition_method, version_introduced, enabled ` +
		`FROM jlabo.wsub_postinstall_messages ` +
		`WHERE postinstall_message_id = ?`

	// run query
	XOLog(sqlstr, postinstallMessageID)
	wpm := WsubPostinstallMessage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, postinstallMessageID).Scan(&wpm.PostinstallMessageID, &wpm.ExtensionID, &wpm.TitleKey, &wpm.DescriptionKey, &wpm.ActionKey, &wpm.LanguageExtension, &wpm.LanguageClientID, &wpm.Type, &wpm.ActionFile, &wpm.Action, &wpm.ConditionFile, &wpm.ConditionMethod, &wpm.VersionIntroduced, &wpm.Enabled)
	if err != nil {
		return nil, err
	}

	return &wpm, nil
}
