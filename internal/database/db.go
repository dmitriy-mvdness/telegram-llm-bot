package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`PRAGMA foreign_keys=ON;`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
