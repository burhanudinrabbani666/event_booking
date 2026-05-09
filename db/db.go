package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	cfg := pq.Config{
		Host:           "localhost",
		Port:           5432,
		User:           "rabbani",
		Password:       "rabbani",
		Database:       "event_bookings",
		ConnectTimeout: 5 * time.Second,
		SSLMode:        pq.SSLModeDisable,
	}

	conn, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		fmt.Println("Failed connect To Database")
		fmt.Println(err)

		return nil, err
	}

	db := sql.OpenDB(conn)

	if err = db.Ping(); err != nil {
		fmt.Println("Failed connect to Database", err)
		return nil, err
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
			id 				SERIAL PRIMARY KEY,
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
			id 					SERIAL PRIMARY KEY,
			name 				VARCHAR(225) NOT NULL,
			description TEXT NOT NULL,
			location 		VARCHAR(255) NOT NULL,
			dateTime 		TIMESTAMP WITH TIME ZONE NOT NULL,
			user_id 		INT REFERENCES users (id) ON DELETE SET NULL,
			createdAt 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updatedAt 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err = db.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Could not Create Events Table", err)
		return err
	}

	return nil
}
