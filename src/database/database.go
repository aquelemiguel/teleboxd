package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./groundhog.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, err
}
