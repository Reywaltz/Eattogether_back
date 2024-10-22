package services

import (
	"eattogether/internal/models"
	"eattogether/internal/repositories"
	env "eattogether/pkg/env"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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

	fmt.Println(tmp.Username, tmp.Password)
	fmt.Printf("Got login payload: %v\n", tmp)

	// TODO db fetch

	payload, err := l.generateJWT(time.Hour, []Roles{ADMIN})
	if err != nil {
		c.String(http.StatusInternalServerError, "Can't generate JWT")
	}

	fmt.Println(payload)

	return c.JSON(http.StatusCreated, &payload)
}

func (l *LoginService) generateJWT(expire time.Duration, roles []Roles) (models.JWTResponse, error) {
	exp := time.Now().Add(expire)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"roles":   roles,
		"user_id": 1,
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
