package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

// InitConnectionDB connection
func InitConnectionDB() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_SCHEMA")
	dbPath := os.Getenv("DB_PATH")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s search_path=%s sslmode=disable TimeZone=Asia/Jakarta", dbUser, dbPass, dbHost, dbPort, dbName, dbPath)

	db = createConnectionPostgres(dsn)

	log.Printf("Successfully connected to database: %s:%s/%s(%s)", dbHost, dbPort, dbName, dbPath)

}

// GetPostgresDB func
func GetPostgresDB() *sql.DB {
	return db
}

func createConnectionPostgres(dsn string) *sql.DB {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	err = checkDefaultTable(db)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}

func checkDefaultTable(db *sql.DB) error {

	queryTodos := `
	CREATE TABLE IF NOT EXISTS tbl_todos(
		todo_id		SERIAL			PRIMARY KEY,
		todo_name	VARCHAR(45)		NOT NULL,
		is_done		BOOL		NOT NULL,
		created_at	TIMESTAMPTZ		NOT NULL,
		updated_at	TIMESTAMPTZ		NOT NULL
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
