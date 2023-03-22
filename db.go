package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// A Entry is a full data struct to be entered into the database
type Entry struct {
	Entry     int    `json:"entry"`
	Sex       string `json:"sex"`
	Activity  string `json:"activity"`
	TimeBegin int    `json:"time_begin"`
	TimeEnd   int    `json:"time_end"`
	TimeTotal int    `json:"time_total"`
	TimeBreak string `json:"time_break"`
	LocBegin  string `json:"location_begin"`
	LocEnd    string `json:"location_end"`
}

// DB is a global variable that contains the sqlite db instance
var DB *sql.DB

// create table if it doesn't exist
const create string = `
CREATE TABLE IF NOT EXISTS entries (
entry INTEGER NOT NULL PRIMARY KEY,
sex TEXT,
activity TEXT,
time_begin INTEGER,
time_end INTEGER,
time_total INTEGER,
time_break TEXT,
location_begin TEXT,
location_end TEXT);`

// connectDatabase opens database if it exists and creates it if it doesn't.
// If the table `entries` does not exist, it creates that as well. Returns
// nil if no error. SIDE-EFFECT - loads database instance into global 
// variable DB.
func connectDatabase() error {
	db, err := sql.Open("sqlite", "./entries.db")
	if err != nil {
		return err
	}
	if _, err := db.Exec(create); err != nil {
		return err
	}
	DB = db
	return nil
}

// runInsert inserts an entry into the `entries` table of the database. Returns
// `true, nil` if successful and `false, error` if not.
func runInsert(oneEntry Entry) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("INSERT INTO entries(entry, sex, activity, time_begin, time_end, time_total, time_break, location_begin, location_end) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(oneEntry.Entry, oneEntry.Sex, oneEntry.Activity,
		oneEntry.TimeBegin, oneEntry.TimeEnd, oneEntry.TimeTotal,
		oneEntry.TimeBreak, oneEntry.LocBegin, oneEntry.LocEnd)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

// runSelect runs a query against the database fetching all entries. Returns
// a slice of all entries or `nil, error` if there is a problem.
func runSelect() ([]Entry, error) {
	rows, err := DB.Query("SELECT * FROM entries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]Entry, 0)

	for rows.Next() {
		oneEntry := Entry{}

		err = rows.Scan(&oneEntry.Entry, &oneEntry.Sex, &oneEntry.Activity,
			&oneEntry.TimeBegin, &oneEntry.TimeEnd, &oneEntry.TimeTotal,
			&oneEntry.TimeBreak, &oneEntry.LocBegin, &oneEntry.LocEnd)

		if err != nil {
			return nil, err
		}
		entries = append(entries, oneEntry)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return entries, err
}
