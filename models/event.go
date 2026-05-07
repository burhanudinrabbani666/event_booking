package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	Datetime    time.Time `binding:"required" json:"datetime"`
	User_id     int       `json:"userId"`
}

type EventCompleteData struct {
	Event
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var events []Event = []Event{}

func (event *Event) Save(db *sql.DB) error {
	// TODO: Add it to database
	query := `
		INSERT INTO events (name, description, location, datetime, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id ;
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		event.Name,
		event.Description,
		event.Location,
		event.Datetime,
		event.User_id,
	).Scan(&event.Id)

	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents(db *sql.DB) ([]EventCompleteData, error) {
	query := `
		SELECT *
		FROM events;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []EventCompleteData
	for rows.Next() {
		var event EventCompleteData
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Datetime,
			&event.User_id,
			&event.CreatedAt,
			&event.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil

}

func GetEventById(db *sql.DB, id int64) (*EventCompleteData, error) {
	query := `
		SELECT *
		FROM events
		WHERE id = $1;
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, errors.New("Error Query.")
	}
	defer stmt.Close()

	var event EventCompleteData
	err = stmt.QueryRow(id).Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.Datetime,
		&event.User_id,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err

	}

	return &event, nil
}
