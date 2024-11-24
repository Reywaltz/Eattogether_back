package models

import "github.com/google/uuid"

type VotePayload struct {
	PlacesIDS []int     `json:"places_ids"`
	RoomID    uuid.UUID `json:"room_id"`
}
type Vote struct {
	RoomID  int `json:"room_id"`
	UserID  int `json:"user_id"`
	PlaceID int `json:"place_id"`
}

type PathRoomID struct {
	RoomID uuid.UUID `param:"roomID"`
}

type UserVotePayload struct {
	Items []Vote `json:"items"`
}

type VoteResultItem struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type VoteResult struct {
	Items []VoteResultItem `json:"items"`
}
