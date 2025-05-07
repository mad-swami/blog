package database

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const DbFile = "blog.db"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DbFile)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func SetupDB() (*sql.DB, error) {
	os.Remove(DbFile)

	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	schemaPath := filepath.Join("database", "schema.sql")
	schema, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	return db, nil
}
