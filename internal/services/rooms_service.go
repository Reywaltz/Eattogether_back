package services

import (
	"eattogether/internal/additions"
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	"eattogether/pkg/customerrors"
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

func (r *RoomsService) GetRoom(c echo.Context) error {
	userID := c.Get("user_id").(int)

	if userID == 0 {
		fmt.Println("No user id")
		return c.String(http.StatusUnauthorized, "No user id")
	}

	roomID, err := uuid.Parse(c.Param("roomID"))
	if err != nil {
		fmt.Println("Can't decode uuid from path", err)
		c.String(http.StatusNotFound, "Not found")
		return nil
	}

	room, err := r.RoomsRepo.GetRoom(roomID)
	if err != nil {
		fmt.Println("Failed to get room", err)
		c.String(http.StatusInternalServerError, "Can't create room")
	}

	return c.JSON(http.StatusOK, &room)

}

func (r *RoomsService) CreateRoom(c echo.Context) error {
	var roomPayload models.RoomCreatePayload

	userID, err := additions.RetriveUserAndPayload(c, &roomPayload, false)
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

	externalID := uuid.New()

	err = r.RoomsRepo.CreateRoom(roomPayload.Name, externalID, userID)
	if err != nil {
		fmt.Println("Failed to create room", err)
		c.String(http.StatusInternalServerError, "Can't create room")
	}

	c.String(http.StatusCreated, "OK")
	return nil
}

func (r *RoomsService) DeleteRoom(c echo.Context) error {
	userID := c.Get("user_id").(int)

	roomID, err := uuid.Parse(c.Param("roomID"))
	if err != nil {
		fmt.Println("Can't decode uuid from path", err)
		c.String(http.StatusNotFound, "Not found")
		return nil
	}

	err = r.RoomsRepo.DeleteRoom(roomID, userID)
	if err != nil {
		fmt.Println("Failed to delete room", err)
		c.String(http.StatusInternalServerError, "Can't delete room")
	}

	c.String(http.StatusCreated, "OK")
	return nil
}

func (r *RoomsService) UpdateRoom(c echo.Context) error {
	userID := c.Get("user_id").(int)

	roomID, err := uuid.Parse(c.Param("roomID"))
	if err != nil {
		fmt.Println("Can't decode uuid from path", err)
	}

	var updatePayload models.RoomUpdatePayload

	if err := c.Bind(&updatePayload); err != nil {
		fmt.Println("Can't bind payload", err)
	}

	fmt.Println(userID, roomID, updatePayload)

	err = r.RoomsRepo.UpdateRoom(roomID, userID, updatePayload.Name)
	if err != nil {
		fmt.Println("Failed to update room", err)
		c.String(http.StatusInternalServerError, "Can't update room")
	}

	c.String(http.StatusNoContent, "OK")
	return nil
}

func CreateRoomsService(rooms_repository *repositories.RoomsRepo) (*RoomsService, error) {
	return &RoomsService{
		RoomsRepo: rooms_repository,
	}, nil
}
