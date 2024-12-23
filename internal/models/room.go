package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID uuid.UUID `json:"external_id"`
	OwnerID    int       `json:"owner_id"`
}

type RoomPayload struct {
	Items []Room `json:"items"`
}

type RoomUpdatePayload struct {
	Name string `json:"name"`
}

type RoomCreatePayload struct {
	Name string `json:"name"`
}
