package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// open connection
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	// connection pooling configuration
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
