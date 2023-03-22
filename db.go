package main

import (
	"database/sql"
	
	_ "modernc.org/sqlite"
)

type Entry struct {
	Entry int `json:"entry"`
	Sex string `json:"sex"`
	Activity string `json:"activity"`
	TimeBegin int `json:"time_begin"`
	TimeEnd int `json:"time_end"`
	TimeTotal int `json:"time_total"`
	TimeBreak string `json:"time_break"`
	LocBegin string `json:"location_begin"`
	LocEnd string `json:"location_end"`
}

var DB *sql.DB

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

func runInsert(oneEntry Entry) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("INSERT INTO entries(entry, sex, activity, time_begin, time_end, time_total, time_break, location_begin, location_end) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(oneEntry.Entry, oneEntry.Sex, oneEntry.Activity,
		oneEntry.TimeBegin, oneEntry.TimeEnd, oneEntry.TimeTotal,
		oneEntry.TimeBreak, oneEntry.LocBegin, oneEntry.LocEnd)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

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
