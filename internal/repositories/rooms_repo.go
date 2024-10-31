package repositories

import (
	"context"
	db "eattogether/pkg/db"

	"eattogether/internal/models"
)

type RoomsRepo struct {
	DB db.PgxInterface
}

func (r *RoomsRepo) GetRooms() ([]models.Room, error) {
	var rooms []models.Room
	res, err := r.DB.Query(context.Background(), "SELECT * FROM rooms")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var room models.Room
		res.Scan(&room.ID, &room.Name)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func CreateRoomsRepo(db *db.DB) *RoomsRepo {
	return &RoomsRepo{
		DB: db,
	}
}
