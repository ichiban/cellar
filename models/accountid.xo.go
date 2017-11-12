// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

// AccountID represents a row from '[custom account_id]'.
type AccountID struct {
	ID int // id
}

// GetAccountIDs runs a custom query, returning results as AccountID.
func GetAccountIDs(db XODB) ([]*AccountID, error) {
	var err error

	// sql query
	const sqlstr = `SELECT id FROM accounts ORDER BY id`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AccountID{}
	for q.Next() {
		a := AccountID{}

		// scan
		err = q.Scan(&a.ID)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}
