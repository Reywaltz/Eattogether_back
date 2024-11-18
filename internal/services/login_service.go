package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type LoginService struct {
	UserRepo  *repositories.UserRepo
	JwtSecret []byte
}

type Roles string

const (
	ADMIN Roles = "admin"
	USER  Roles = "user"
)

func (l *LoginService) LoginHandler(c echo.Context) error {
	var tmp models.LoginPayload

	err := c.Bind(&tmp)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad request")
	}

	row, err := l.UserRepo.GetUser(tmp.Username, tmp.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No user", err)
			return c.String(http.StatusBadRequest, "No user")
		} else {
			fmt.Println("DB error", err)
			return c.String(http.StatusBadGateway, "DB error")
		}
	}

	payload, err := l.generateJWT(time.Hour, row)
	if err != nil {
		c.String(http.StatusInternalServerError, "Can't generate JWT")
	}

	return c.JSON(http.StatusCreated, &payload)
}

func (l *LoginService) generateJWT(expire time.Duration, user models.User) (models.JWTResponse, error) {
	exp := time.Now().Add(expire)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"roles":   []string{user.Role},
		"user_id": user.ID,
		"exp":     exp.Unix(),
		"iat":     time.Now().UTC().Unix(),
	})

	env_reader, _ := env.GetENVReader()

	tokenString, err := claims.SignedString(env_reader.JWT_SECRET)
	if err != nil {
		fmt.Printf("Error during jwt creation: %v", err)
		return models.JWTResponse{}, err
	}

	// TODO write token to db

	return models.JWTResponse{
		Token:  tokenString,
		Expire: exp.Unix(),
	}, nil
}

func CreateLoginService(
	user_repository *repositories.UserRepo,
) (*LoginService, error) {
	return &LoginService{
		UserRepo: user_repository,
	}, nil
}
