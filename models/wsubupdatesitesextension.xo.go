// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubUpdateSitesExtension represents a row from 'jlabo.wsub_update_sites_extensions'.
type WsubUpdateSitesExtension struct {
	UpdateSiteID int `json:"update_site_id"` // update_site_id
	ExtensionID  int `json:"extension_id"`   // extension_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubUpdateSitesExtension exists in the database.
func (wuse *WsubUpdateSitesExtension) Exists() bool {
	return wuse._exists
}

// Deleted provides information if the WsubUpdateSitesExtension has been deleted from the database.
func (wuse *WsubUpdateSitesExtension) Deleted() bool {
	return wuse._deleted
}

// Insert inserts the WsubUpdateSitesExtension to the database.
func (wuse *WsubUpdateSitesExtension) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wuse._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO jlabo.wsub_update_sites_extensions (` +
		`update_site_id, extension_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wuse.UpdateSiteID, wuse.ExtensionID)
	_, err = db.Exec(sqlstr, wuse.UpdateSiteID, wuse.ExtensionID)
	if err != nil {
		return err
	}

	// set existence
	wuse._exists = true

	return nil
}

// Update statements omitted due to lack of fields other than primary key

// Delete deletes the WsubUpdateSitesExtension from the database.
func (wuse *WsubUpdateSitesExtension) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wuse._exists {
		return nil
	}

	// if deleted, bail
	if wuse._deleted {
		return nil
	}

	// sql query with composite primary key
	const sqlstr = `DELETE FROM jlabo.wsub_update_sites_extensions WHERE update_site_id = ? AND extension_id = ?`

	// run query
	XOLog(sqlstr, wuse.UpdateSiteID, wuse.ExtensionID)
	_, err = db.Exec(sqlstr, wuse.UpdateSiteID, wuse.ExtensionID)
	if err != nil {
		return err
	}

	// set deleted
	wuse._deleted = true

	return nil
}

// WsubUpdateSitesExtensionByExtensionID retrieves a row from 'jlabo.wsub_update_sites_extensions' as a WsubUpdateSitesExtension.
//
// Generated from index 'wsub_update_sites_extensions_extension_id_pkey'.
func WsubUpdateSitesExtensionByExtensionID(db XODB, extensionID int) (*WsubUpdateSitesExtension, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`update_site_id, extension_id ` +
		`FROM jlabo.wsub_update_sites_extensions ` +
		`WHERE extension_id = ?`

	// run query
	XOLog(sqlstr, extensionID)
	wuse := WsubUpdateSitesExtension{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, extensionID).Scan(&wuse.UpdateSiteID, &wuse.ExtensionID)
	if err != nil {
		return nil, err
	}

	return &wuse, nil
}
