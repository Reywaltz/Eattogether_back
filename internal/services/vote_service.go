package services

import (
	"eattogether/internal/additions"
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VotesService struct {
	VotesRepo *repositories.VotesRepo
	RoomsRepo *repositories.RoomsRepo
	ENVReader *env.EnvReader
}

func (v *VotesService) GetUserVotes(c echo.Context) error {
	var request models.PathRoomID

	userID, err := additions.RetriveUserAndPayload(c, &request, false)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.JSONMessage{
			Message: "Bad payload",
		})
	}

	room, err := v.RoomsRepo.GetRoom(request.RoomID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.JSONMessage{
			Message: "room not found",
		})
	}

	votes, err := v.VotesRepo.GetVotesByUser(userID, room.ID)
	if err != nil {
		fmt.Printf("Can't get data from db:%v\n", err)
	}

	var response models.UserVotePayload

	if votes != nil {
		response.Items = votes
	} else {
		response.Items = []models.Vote{}
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (v *VotesService) GetVotingResult(c echo.Context) error {
	var roomID models.PathRoomID

	_, err := additions.RetriveUserAndPayload(c, &roomID, false)
	if err != nil {
		fmt.Println("No user id")
		return c.JSON(http.StatusBadRequest, models.JSONMessage{
			Message: "Bad request",
		})
	}

	room, err := v.RoomsRepo.GetRoom(roomID.RoomID)
	if err != nil {
		fmt.Printf("No such room: %v\n", err)
		return c.JSON(http.StatusNotFound, models.JSONMessage{
			Message: "No such room",
		})
	}

	result, _ := v.VotesRepo.GetVotingResultByRoom(room.ID)

	var response models.VoteResult
	if result != nil {
		response.Items = result
	} else {
		response.Items = []models.VoteResultItem{}
	}

	return c.JSON(http.StatusOK, response)
}

func CreateVotesService(
	vote_repository *repositories.VotesRepo,
	room_repository *repositories.RoomsRepo,
) (*VotesService, error) {
	return &VotesService{
		VotesRepo: vote_repository,
		RoomsRepo: room_repository,
	}, nil
}
