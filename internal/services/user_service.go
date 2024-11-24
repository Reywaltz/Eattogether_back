package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UsersService struct {
	UserRepo  *repositories.UserRepo
	RoomsRepo *repositories.RoomsRepo
}

func (u *UsersService) GetUsersByRoom(c echo.Context) error {
	roomUUID, err := uuid.Parse(c.QueryParam("room_id"))
	if err != nil {
		fmt.Println("Can't decode uuid from path", err)
	}

	room, err := u.RoomsRepo.GetRoom(roomUUID)
	if err != nil {
		fmt.Printf("Room not found: %v\n", err)
	}

	users, err := u.RoomsRepo.GetUsersByRoom(room.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.JSONMessage{
			Message: "Can't get users",
		})
	}
	return c.JSON(http.StatusOK, users)

}

func CreateUsersService(user_repo *repositories.UserRepo, rooms_repo *repositories.RoomsRepo) (*UsersService, error) {
	return &UsersService{
		UserRepo:  user_repo,
		RoomsRepo: rooms_repo,
	}, nil
}
