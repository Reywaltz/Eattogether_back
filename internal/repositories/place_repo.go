package repositories

import (
	"context"
	db "eattogether/pkg/db"

	"eattogether/internal/models"
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

	res.Scan(&places)
	return places, nil
}

func CreatePlaceRepo(db *db.DB) *PlacesRepo {
	return &PlacesRepo{
		DB: db,
	}
}
