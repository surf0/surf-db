package models

type Record struct {
	Id string `json:"id"`
	Timestamp string `json:"timestamp"`
	PlayerName string `json:"player_name"`
	PlayerId string `json:"player_id"`
	Type string `json:"type"`
	Track int `json:"track"`
	MapName string `json:"map_name"`
	Time string `json:"time"`
	Improvement string `json:"improvement"`
	Server string `json:"server"`
}