package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Database struct {}

func (database Database) Setup() {
	config := newDatabaseConfig()

    db, err := sql.Open("mysql", config.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	defer db.Close()
}