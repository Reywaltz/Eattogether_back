package services

import (
	"eattogether/internal/additions"
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	"eattogether/pkg/customerrors"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlaceService struct {
	PlaceRepo *repositories.PlacesRepo
	RoomRepo  *repositories.RoomsRepo
	ENVReader *env.EnvReader
}

func (p *PlaceService) GetPlaces(c echo.Context) error {
	var response models.PlacePayload

	places, err := p.PlaceRepo.GetPlaces()
	if err != nil {
		fmt.Printf("Can't get data from db:%v\n", err)
	}

	if places != nil {
		response.Items = places
	} else {
		response.Items = []models.Place{}
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (p *PlaceService) Vote(c echo.Context) error {
	var body models.VotePayload

	userID, err := additions.RetriveUserAndPayload(c, &body, false)
	if err != nil {
		switch err.(type) {
		case *customerrors.DataNotBindable:
			return c.JSON(http.StatusBadRequest, models.JSONMessage{
				Message: "wrong payload",
			})

		case *customerrors.UserNotSetError:
			return c.JSON(http.StatusUnauthorized, models.JSONMessage{
				Message: "User not set",
			})
		}
	}

	room, err := p.RoomRepo.GetRoom(body.RoomID)
	if err != nil {
		fmt.Printf("error during room get: %v\n", err)
	}

	err = p.PlaceRepo.InsertVotes(room.ID, userID, body.PlacesIDS)
	if err != nil {
		fmt.Printf("Can't vote: %v\n", err)
	}

	c.String(http.StatusOK, "VOTED!!!!")
	return nil
}

func CreatePlacesService(place_repository *repositories.PlacesRepo, room_repository *repositories.RoomsRepo) (*PlaceService, error) {
	return &PlaceService{
		PlaceRepo: place_repository,
		RoomRepo:  room_repository,
	}, nil
}
