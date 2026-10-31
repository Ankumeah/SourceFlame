package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DBController struct{ db *sql.DB }

func GetDBController(url string) (*DBController, error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DBController{db}, nil
}

// This exposes the raw underlying DB object.
// This is only to be used to execute migrations
func (db *DBController) DB() *sql.DB {
	return db.db
}
