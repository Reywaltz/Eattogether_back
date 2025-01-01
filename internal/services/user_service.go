package services

import (
	"eattogether/internal/additions"
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	"eattogether/pkg/elastic"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UsersService struct {
	UserRepo      *repositories.UserRepo
	RoomsRepo     *repositories.RoomsRepo
	ElasticClient *elastic.ElasticClient
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

func (u *UsersService) SearchUsers(c echo.Context) error {
	search_param := c.QueryParam("user_search")
	query := fmt.Sprintf(
		`{"query": {"match_phrase_prefix": {"name": {"query": "%s"}}}}`,
		search_param,
	)

	result, err := u.ElasticClient.Search("users", query)
	if err != nil {
		fmt.Printf("Error during elastic call: %v\n", err)
		c.JSON(http.StatusBadGateway, models.JSONMessage{
			Message: "Internal Error",
		})
		return err
	}

	var user models.ElasticUser

	results := additions.ParseElasticResult(result, user)
	if len(results) == 0 {
		c.JSON(http.StatusOK, []models.ElasticUser{})
		return nil
	}

	c.JSON(http.StatusOK, results)
	return nil
}

func CreateUsersService(
	user_repo *repositories.UserRepo,
	rooms_repo *repositories.RoomsRepo,
	elastic_client *elastic.ElasticClient,
) (*UsersService, error) {
	return &UsersService{
		UserRepo:      user_repo,
		RoomsRepo:     rooms_repo,
		ElasticClient: elastic_client,
	}, nil
}
