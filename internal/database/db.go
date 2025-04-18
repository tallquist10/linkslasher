package database

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func SqlLite(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pass@/links")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func New() (*sql.DB, error) {
	sqlMode := os.Getenv("SQL_MODE")
	if sqlMode == "MYSQL" {
		return MySql()
	}
	if sqlMode == "SQLITE" {
		return SqlLite("linkslasher.db")
	}
	return nil, errors.New("Only MYSQL and SQLITE are valid options for SQL_MODE")
}
