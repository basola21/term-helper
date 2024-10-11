package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "responses.db"

var db *sql.DB

// Initialize the database connection and create the responses table if it doesn't exist.
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS responses (
		id TEXT PRIMARY KEY,
		model TEXT,
		message TEXT,
		created REAL,
		usage_tokens INTEGER,
		usage_time REAL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

// SaveResponse saves the API response into the database.
func SaveResponse(id, model, message string, created float64, tokens int, time float64) error {
	insertQuery := `INSERT INTO responses (id, model, message, created, usage_tokens, usage_time) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(insertQuery, id, model, message, created, tokens, time)
	if err != nil {
		return fmt.Errorf("error inserting data into DB: %v", err)
	}
	return nil
}

type usedTokens struct {
	id     string
	tokens int
}

func GetUserTokens() (int, error) {
	query := `SELECT SUM(usage_tokens) AS total_tokens FROM responses`

	var result int
	err := db.QueryRow(query).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, fmt.Errorf("you do not seem to have any tokens")
		}
		return result, fmt.Errorf("error querying data from DB: %v", err)
	}

	return result, nil
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
