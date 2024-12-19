package models

import (
	"time"

	"github.com/ThomasPhilipp/go-event-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{}

// method
func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Exec() changes data (insert, update, etc.)
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // return auto generated id
	e.ID = id

	return err
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

// function
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// Query() fetches data
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var events []Event
	for result.Next() {
		var event Event
		result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	// Query() fetches data
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
