package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// function to return db connection and error if there are any
func openDB(dsn, driver string) (*sql.DB, error) {

	// Use driver that is appliable to DB you happen to use/prefer
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	// ping db to test connectivity
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
