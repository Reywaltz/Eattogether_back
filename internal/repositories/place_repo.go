package repositories

import (
	"context"
	db "eattogether/pkg/db"
	"fmt"

	"eattogether/internal/models"

	"github.com/jackc/pgx/v5"
)

type PlacesRepo struct {
	DB db.PgxInterface
}

func (p *PlacesRepo) GetPlaces() ([]models.Place, error) {
	var places []models.Place
	res, err := p.DB.Query(context.Background(), "SELECT * FROM places")
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var place models.Place
		err = res.Scan(&place.ID, &place.Name, &place.Image)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

func (p *PlacesRepo) InsertVotes(roomID int, userID int, placeIDS []int) error {
	var rows [][]interface{}

	// Собираем данные для вставки одним запросом
	for _, placeID := range placeIDS {
		rows = append(rows, []interface{}{roomID, userID, placeID})
	}

	transaction, err := p.DB.Begin(context.Background())
	if err != nil {
		fmt.Printf("Failed to begin transaction %v\n", err)
		return err
	}

	defer transaction.Rollback(context.Background())

	_, err = transaction.CopyFrom(
		context.Background(),
		pgx.Identifier{"place_votes"},
		[]string{"room_id", "user_id", "place_id"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		fmt.Printf("Failed to copy data: %v\n", err)
		return err
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		fmt.Printf("Failed to commit: %v\n", err)
		return err
	}

	return nil
}

func CreatePlaceRepo(db *db.DB) *PlacesRepo {
	return &PlacesRepo{
		DB: db,
	}
}
