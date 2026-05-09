package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://rabbani:rabbani@localhost:5432/event_bookings?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to open Database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := CreateTable(db); err != nil {
		return nil, err
	}

	return db, nil

}

func CreateTable(db *sql.DB) error {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users(
			id 			SERIAL PRIMARY KEY,
			email 		TEXT NOT NULL UNIQUE,
			password	TEXT NOT NULL
		)
	`
	_, err := db.Exec(createUserTable)
	if err != nil {
		fmt.Println("Could not Create Users Table", err)
		return err
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events(
			id 				SERIAL PRIMARY KEY,
			name 			VARCHAR(225) NOT NULL,
			description 	TEXT NOT NULL,
			location 		VARCHAR(255) NOT NULL,
			dateTime 		TIMESTAMP WITH TIME ZONE NOT NULL,
			user_id 		INT REFERENCES users (id) ON DELETE SET NULL,
			createdAt 		TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updatedAt 		TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err = db.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Could not Create Events Table", err)
		return err
	}

	creatRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations(
			id 			SERIAL PRIMARY KEY,
			event_id	INT REFERENCES events(id),  
			user_id 	INT REFERENCES users(id), 
			createdAt 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updatedAt 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err = db.Exec(creatRegistrationsTable)
	if err != nil {
		fmt.Println("Could not Create Registations Table", err)
		return err
	}

	return nil
}
