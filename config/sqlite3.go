package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitConnectionDB connection
func InitConnectionDB() {
	db = createConnectionSQLite3()
}

var db *sql.DB

func GetSQLite3Conn() *sql.DB {
	return db
}

func createConnectionSQLite3() *sql.DB {

	db, err := sql.Open("sqlite3", "./todos.db")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	err = checkDefaultTable(db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func checkDefaultTable(db *sql.DB) error {

	queryTodos := `
	CREATE TABLE IF NOT EXISTS tbl_todos(
		todo_id		INTEGER			PRIMARY KEY,
		todo_name	VARCHAR(45)		NOT NULL,
		is_done		TINYINT(1)		NOT NULL,
		created_at	TIMESTAMP		NOT NULL,
		updated_at	TIMESTAMP		NOT NULL
	)
	`

	statement, err := db.Prepare(queryTodos)

	if err != nil {
		return err
	}

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	return nil
}
