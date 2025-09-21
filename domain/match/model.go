package match

import "time"

type Match struct {
	EventId  string    `json:"event_id" bson:"event_id"`
	CourtNo  int       `json:"court_no" bson:"court_no"`
	DateTime time.Time `json:"date_time" bson:"date_time"`
	Status   string    `json:"status" bson:"status"`
	TeamA    Teamate   `json:"team_a" bson:"team_a"`
	TeamB    Teamate   `json:"team_b" bson:"team_b"`
	WhoWon   string    `json:"who_won" bson:"who_won"`
}

type Teamate struct {
	Player1 string `json:"player_1" bson:"player_1"`
	Player2 string `json:"player_2" bson:"player_2"`
}
