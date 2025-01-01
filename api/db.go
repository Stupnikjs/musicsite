package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// connect db
func ConnectDB(uri string) (*sql.DB, error) {
	db, err := openDB(uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openDB(uri string) (*sql.DB, error) {
	// Open the database connection
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatalf("Unable to open connection: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		// log.Fatalf("Unable to connect to database: %v", err)
		fmt.Println("no connection to db")
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, err
}

// open db
