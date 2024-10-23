package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlaceService struct {
	PlaceRepo *repositories.PlacesRepo
	ENVReader *env.EnvReader
}

func (p *PlaceService) GetPlaces(c echo.Context) error {
	var response models.PlacePayload

	places, err := p.PlaceRepo.GetPlaces()
	if err != nil {
		fmt.Printf("Can't get data from db:%v\n", err)
	}

	response.Items = places

	c.JSON(http.StatusOK, response)
	return nil
}

func CreatePlacesService(place_repository *repositories.PlacesRepo) (*PlaceService, error) {
	return &PlaceService{
		PlaceRepo: place_repository,
	}, nil
}
