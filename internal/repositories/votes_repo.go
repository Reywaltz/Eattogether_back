package repositories

import (
	"context"
	db "eattogether/pkg/db"
	"fmt"

	"eattogether/internal/models"
)

type VotesRepo struct {
	DB db.PgxInterface
}

func (v *VotesRepo) GetVotesByUser(userID int, roomID int) ([]models.Vote, error) {
	var votes []models.Vote
	res, err := v.DB.Query(
		context.Background(),
		"SELECT room_id, user_id, place_id FROM place_votes WHERE user_id=$1 and room_id=$2",
		userID,
		roomID,
	)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var vote models.Vote
		err = res.Scan(&vote.RoomID, &vote.UserID, &vote.PlaceID)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	return votes, nil
}

func (v *VotesRepo) GetVotingResultByRoom(roomID int) ([]models.VoteResultItem, error) {

	res, err := v.DB.Query(
		context.Background(),
		`SELECT 
		places.name,
		COUNT(place_votes.place_id)
		FROM place_votes
		INNER JOIN places ON places.id = place_votes.place_id
		WHERE place_votes.room_id=$1
		GROUP BY places.name`,
		roomID,
	)

	if err != nil {
		fmt.Printf("error during query: %v\n", err)
		return nil, err
	}

	defer res.Close()

	var result []models.VoteResultItem
	for res.Next() {
		var votingResult models.VoteResultItem
		err := res.Scan(&votingResult.Name, &votingResult.Count)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
			return nil, err
		}

		result = append(result, votingResult)
	}

	return result, nil
}

func CreateVotesRepo(db *db.DB) *VotesRepo {
	return &VotesRepo{
		DB: db,
	}
}
