package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Global variable for database connection
var db *sql.DB

// Function for connecting to database
func ConnectDB() {
	var err error

	// Get database URL from environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Function for closing database connection
func CloseDB() {
	db.Close()
}

// Function for executing SQL query
func ExecSQL(query string, args ...interface{}) (sql.Result, error) {
	if db == nil {
		ConnectDB()
	}

	res, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Function for executing SQL query and returning rows
func QuerySQL(query string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		ConnectDB()
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Function for executing SQL query and returning single row
func QueryRowSQL(query string, args ...interface{}) *sql.Row {
	if db == nil {
		ConnectDB()
	}

	row := db.QueryRow(query, args...)

	return row
}

// Function for executing SQL query and returning single value
func QueryValueSQL(query string, args ...interface{}) (interface{}, error) {
	if db == nil {
		ConnectDB()
	}

	row := db.QueryRow(query, args...)
	var value interface{}
	err := row.Scan(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
