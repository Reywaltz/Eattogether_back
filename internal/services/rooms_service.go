package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RoomsService struct {
	RoomsRepo *repositories.RoomsRepo
	ENVReader *env.EnvReader
}

func (r *RoomsService) GetRooms(c echo.Context) error {
	userID := c.Get("user_id").(int)
	var response models.RoomPayload

	rooms, err := r.RoomsRepo.GetRooms(userID)
	if err != nil {
		fmt.Printf("Can't get data from db:%v\n", err)
	}

	response.Items = rooms

	c.JSON(http.StatusOK, response)
	return nil
}

func (r *RoomsService) CreateRoom(c echo.Context) error {
	user_id := c.Get("user_id").(int)

	var roomPayload models.RoomCreatePayload
	err := c.Bind(&roomPayload)
	if err != nil {
		fmt.Println("Can't bind", err)
	}

	external_id := uuid.New()

	err = r.RoomsRepo.CreateRoom(roomPayload.Name, external_id, user_id)
	if err != nil {
		fmt.Println("Failed to create room", err)
		c.String(http.StatusInternalServerError, "Can't create room")
	}

	c.String(http.StatusCreated, "OK")
	return nil
}

func CreateRoomsService(rooms_repository *repositories.RoomsRepo) (*RoomsService, error) {
	return &RoomsService{
		RoomsRepo: rooms_repository,
	}, nil
}
