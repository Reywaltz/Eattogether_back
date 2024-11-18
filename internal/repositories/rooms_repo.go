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
		"SELECT id, name, created_at, external_id, owner_id FROM rooms where owner_id=$1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var room models.Room
		err := res.Scan(&room.ID, &room.Name, &room.CreatedAt, &room.ExternalID, &room.OwnerID)
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

func (r *RoomsRepo) DeleteRoom(roomID uuid.UUID, userID int) error {
	_, err := r.DB.Exec(
		context.Background(),
		"DELETE FROM rooms WHERE external_id=$1 and owner_id=$2",
		roomID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RoomsRepo) UpdateRoom(roomID uuid.UUID, userID int, name string) error {
	_, err := r.DB.Exec(
		context.Background(),
		"UPDATE rooms SET name=$1 WHERE external_id=$2 and owner_id=$3",
		name,
		roomID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RoomsRepo) GetRoom(roomID uuid.UUID) (models.Room, error) {
	var room models.Room

	row := r.DB.QueryRow(
		context.Background(),
		"SELECT id, name, created_at, external_id, owner_id FROM rooms WHERE external_id=$1",
		roomID,
	)

	err := row.Scan(&room.ID, &room.Name, &room.CreatedAt, &room.ExternalID, &room.OwnerID)
	if err != nil {
		return room, err
	}

	return room, nil
}

func CreateRoomsRepo(db *db.DB) *RoomsRepo {
	return &RoomsRepo{
		DB: db,
	}
}
