package db

import (
	"database/sql"
	"os"
)

const schema = `CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL CHECK(length(title) <= 120),
    comment TEXT NOT NULL CHECK(length(comment) <= 120),
    repeat TEXT NOT NULL CHECK(length(repeat) <= 120)
);

CREATE INDEX idx_scheduler_date
ON scheduler(date);`

var DB *sql.DB

func Init(dbfile string) error {

	dbFile := "scheduler.db"
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	DB, err = sql.Open("sqlite", dbfile)

	if err != nil {
		return err
	}

	if install {
		_, err = DB.Exec(schema)

		if err != nil {
			DB.Close()
			DB = nil
			return err
		}
	}

	return nil
}
