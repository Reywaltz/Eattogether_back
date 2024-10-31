package repositories

import (
	"context"
	db "eattogether/pkg/db"
	"fmt"

	"eattogether/internal/models"

	"github.com/google/uuid"
)

type RoomsRepo struct {
	DB db.PgxInterface
}

func (r *RoomsRepo) GetRooms(userID int) ([]models.Room, error) {
	var rooms []models.Room
	res, err := r.DB.Query(
		context.Background(),
		"SELECT id, name, created_at, external_id FROM rooms where owner_id=$1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var room models.Room
		err := res.Scan(&room.ID, &room.Name, &room.CreatedAt, &room.ExternalID)
		if err != nil {
			fmt.Println(err)
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *RoomsRepo) CreateRoom(name string, externalID uuid.UUID, ownerID int) error {
	_, err := r.DB.Exec(
		context.Background(),
		"INSERT INTO rooms (name, external_id, owner_id) VALUES ($1, $2, $3)",
		name,
		externalID,
		ownerID,
	)

	if err != nil {
		return err
	}

	return nil
}

func CreateRoomsRepo(db *db.DB) *RoomsRepo {
	return &RoomsRepo{
		DB: db,
	}
}
