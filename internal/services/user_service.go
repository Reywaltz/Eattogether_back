package services

import (
	"eattogether/internal/repositories"

	"github.com/labstack/echo/v4"
)

type UsersService struct {
	UserRepo *repositories.UserRepo
}

func (u *UsersService) GetUser(c echo.Context) error {
	return nil
}

func CreateUsersService(place_repository *repositories.UserRepo) (*UsersService, error) {
	return &UsersService{
		UserRepo: place_repository,
	}, nil
}
