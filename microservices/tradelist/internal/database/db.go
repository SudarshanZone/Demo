package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "cdac"
	dbname   = "IRRA"
)

var db *sql.DB

// Init initializes the database connection
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL version:", version)
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	if db == nil {
		log.Println("Warning: GetDB called but db is nil")
	}
	return db
}
