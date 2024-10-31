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

func CreatePlaceRepo(db *db.DB) *PlacesRepo {
	return &PlacesRepo{
		DB: db,
	}
}
