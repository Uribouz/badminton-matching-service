package event

import "time"

type Event struct {
	EventId   string    `json:"event_id" bson:"event_id"`
	EventDate time.Time `json:"event_date" bson:"event_date"`
}