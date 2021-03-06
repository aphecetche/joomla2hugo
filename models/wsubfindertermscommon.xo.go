// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

// WsubFinderTermsCommon represents a row from 'jlabo.wsub_finder_terms_common'.
type WsubFinderTermsCommon struct {
	Term     string `json:"term"`     // term
	Language string `json:"language"` // language
}

// WsubFinderTermsCommonsByLanguage retrieves a row from 'jlabo.wsub_finder_terms_common' as a WsubFinderTermsCommon.
//
// Generated from index 'idx_lang'.
func WsubFinderTermsCommonsByLanguage(db XODB, language string) ([]*WsubFinderTermsCommon, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term, language ` +
		`FROM jlabo.wsub_finder_terms_common ` +
		`WHERE language = ?`

	// run query
	XOLog(sqlstr, language)
	q, err := db.Query(sqlstr, language)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTermsCommon{}
	for q.Next() {
		wftc := WsubFinderTermsCommon{}

		// scan
		err = q.Scan(&wftc.Term, &wftc.Language)
		if err != nil {
			return nil, err
		}

		res = append(res, &wftc)
	}

	return res, nil
}

// WsubFinderTermsCommonsByTermLanguage retrieves a row from 'jlabo.wsub_finder_terms_common' as a WsubFinderTermsCommon.
//
// Generated from index 'idx_word_lang'.
func WsubFinderTermsCommonsByTermLanguage(db XODB, term string, language string) ([]*WsubFinderTermsCommon, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`term, language ` +
		`FROM jlabo.wsub_finder_terms_common ` +
		`WHERE term = ? AND language = ?`

	// run query
	XOLog(sqlstr, term, language)
	q, err := db.Query(sqlstr, term, language)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTermsCommon{}
	for q.Next() {
		wftc := WsubFinderTermsCommon{}

		// scan
		err = q.Scan(&wftc.Term, &wftc.Language)
		if err != nil {
			return nil, err
		}

		res = append(res, &wftc)
	}

	return res, nil
}
