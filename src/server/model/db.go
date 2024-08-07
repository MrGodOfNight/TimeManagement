/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

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

	// Open database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping database connection
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

	// Execute query
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

	// Execute query and return rows
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

	// Execute query and return only one row
	row := db.QueryRow(query, args...)

	return row
}

// Function for executing SQL query and returning single value
func QueryValueSQL(query string, args ...interface{}) (interface{}, error) {
	if db == nil {
		ConnectDB()
	}

	// Execute query and return only one row
	row := db.QueryRow(query, args...)
	// Get value from row
	var value interface{}
	err := row.Scan(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
