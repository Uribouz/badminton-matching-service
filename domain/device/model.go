package device

import "time"

type Device struct {
	Id          string     `json:"id" bson:"id"`
	LastLogin   *time.Time `json:"last_login" bson:"last_login"`
	LastEventId string     `json:"last_event_id" bson:"last_event_id"`
}
