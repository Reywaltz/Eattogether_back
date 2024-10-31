package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomsService struct {
	RoomsRepo *repositories.RoomsRepo
	ENVReader *env.EnvReader
}

func (r *RoomsService) GetRooms(c echo.Context) error {
	var response models.RoomPayload

	rooms, err := r.RoomsRepo.GetRooms()
	if err != nil {
		fmt.Printf("Can't get data from db:%v\n", err)
	}

	response.Items = rooms

	c.JSON(http.StatusOK, response)
	return nil
}

func CreateRoomsService(rooms_repository *repositories.RoomsRepo) (*RoomsService, error) {
	return &RoomsService{
		RoomsRepo: rooms_repository,
	}, nil
}
