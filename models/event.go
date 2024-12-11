package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

// method
func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

// function
func GetAllEvents() []Event {
	return events
}
