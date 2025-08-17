package partner

type Partner struct {
	EventId    string `json:"event_id" bson:"event_id"`
	PlayerName string `json:"player_name" bson:"player_name"`
}