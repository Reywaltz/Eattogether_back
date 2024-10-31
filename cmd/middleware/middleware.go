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
		fmt.Println("MIDDLEWARE")
		cookies, err := c.Cookie("X-Auth-Token")
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, messageNoCookie)
		}

		env_reader, _ := env.GetENVReader()

		var claims models.JWTClaims

		token, err := jwt.ParseWithClaims(cookies.Value, &claims, func(token *jwt.Token) (interface{}, error) {
			return env_reader.JWT_SECRET, nil
		})

		if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
			c.Set("user_id", claims.User_id)
			c.Set("roles", claims.Roles)
		} else {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, messageError)
		}

		return next(c)
	}
}
