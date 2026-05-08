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
	User_id     int       `json:"userId,omitempty"`
}

type EventCompleteData struct {
	Event
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

var events []Event = []Event{}

func (event *Event) Create(DB *sql.DB) error {
	query := `
		INSERT INTO events (name, description, location, datetime, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id ;
	`

	stmt, err := DB.Prepare(query)
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

func GetAllEvents(DB *sql.DB) ([]EventCompleteData, error) {
	query := `
		SELECT *
		FROM events;
	`
	rows, err := DB.Query(query)
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

func GetEventById(DB *sql.DB, id int) (*EventCompleteData, error) {
	query := `
		SELECT *
		FROM events
		WHERE id = $1;
	`

	stmt, err := DB.Prepare(query)
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

func (event EventCompleteData) Update(DB *sql.DB) error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, datetime = $4, updatedat = $5
		WHERE id = $6;
	`

	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.Datetime,
		event.UpdatedAt,
		event.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteEventById(DB *sql.DB, id int) error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil

}
