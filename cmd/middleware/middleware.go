package middleware

import (
	"eattogether/internal/models"
	"eattogether/pkg/env"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	messageError := models.JSONMessage{
		Message: "error",
	}

	messageNoCookie := models.JSONMessage{
		Message: "no cookie",
	}

	return func(c echo.Context) error {
		fmt.Println(c.Cookies())
		cookies, err := c.Cookie("X-Auth-Token")
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, messageNoCookie)
		}

		env_reader, _ := env.GetENVReader()

		_, err = jwt.Parse(cookies.Value, func(token *jwt.Token) (interface{}, error) {
			return env_reader.JWT_SECRET, nil
		})

		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, messageError)
		}

		return next(c)
	}
}
