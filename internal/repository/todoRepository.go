package internal

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Create(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL
	)`)
	return err
}

func Add(db *sql.DB, task string) error {
	_, err := db.Exec("INSERT INTO todos (task) VALUES (?)", task)
	return err
}

func Delete(db *sql.DB, taskID int) (sql.Result, error) {
	res, err := db.Exec("DELETE FROM todos WHERE id = ?", taskID)
	return res, err
}

func List(db *sql.DB) (*sql.Rows, error) {
	res, err := db.Query("SELECT id, task FROM todos")
	return res, err
}
