package models

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoomPayload struct {
	Items []Room `json:"items"`
}
