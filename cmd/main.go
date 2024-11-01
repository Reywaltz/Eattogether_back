package main

import (
	router "eattogether/cmd/handlers"
	repos "eattogether/internal/repositories"
	"eattogether/internal/services"
	db "eattogether/pkg/db"
	"eattogether/pkg/env"
	"eattogether/pkg/logger"
)

func main() {
	logger := logger.CreateLogger()

	envReader, err := env.GetENVReader()
	if err != nil {
		logger.Fatalf("Failed to create envReader: %v", err)
	}

	database, err := db.CreateConnection(envReader.DB_URL)
	if err != nil {
		logger.Fatalf("Failed to create db connection: %v", err)
	}

	logger.Info("Connected to db: %v", database)

	places_repo := repos.CreatePlaceRepo(database)
	user_repo := repos.CreateUserRepo(database)
	rooms_repo := repos.CreateRoomsRepo(database)

	places_service, _ := services.CreatePlacesService(places_repo)
	users_service, _ := services.CreateUsersService(user_repo)
	login_service, _ := services.CreateLoginService(user_repo)
	rooms_service, _ := services.CreateRoomsService(rooms_repo)

	e := router.InitRouter(users_service, places_service, login_service, rooms_service)
	// TODO Correct CORS
	go router.HandleBroadcast()

	e.Start(":8000")

}
