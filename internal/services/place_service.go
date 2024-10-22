package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlaceService struct {
	PlaceRepo *repositories.PlacesRepo
	ENVReader *env.EnvReader
}

func (p *PlaceService) GetPlaces(c echo.Context) error {
	var response models.PlacePayload

	items := make([]models.Place, 10)

	for id := range items {
		items[id].ID = id + 1
		items[id].Name = "Галки"
		items[id].Image = "/Path/to/image"
	}

	response.Items = items

	c.JSON(http.StatusOK, response)
	return nil
}

func CreatePlacesService(place_repository *repositories.PlacesRepo) (*PlaceService, error) {
	return &PlaceService{
		PlaceRepo: place_repository,
	}, nil
}
