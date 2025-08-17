package match

import "time"

type Match struct {
	EventId  string    `json:"event_id" bson:"event_id"`
	CourtNo  int       `json:"court_no" bson:"court_no"`
	DateTime time.Time `json:"date_time" bson:"date_time"`
}