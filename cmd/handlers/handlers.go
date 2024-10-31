package handlers

import (
	custom_middleware "eattogether/cmd/middleware"
	"eattogether/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(
	user_service *services.UsersService,
	place_service *services.PlaceService,
	login_service *services.LoginService,
	rooms_service *services.RoomsService,
) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))

	api_group := e.Group("/api/v1")

	api_group.POST("/login", login_service.LoginHandler)
	api_group.GET("/ws", HandleConnections)

	places_group := api_group.Group("/places", custom_middleware.JWTMiddleware)
	places_group.GET("", place_service.GetPlaces)
	// TODO move to rooms group
	places_group.POST("/vote", place_service.Vote)

	rooms_group := api_group.Group("/rooms", custom_middleware.JWTMiddleware)
	rooms_group.POST("", rooms_service.CreateRoom)
	rooms_group.GET("", rooms_service.GetRooms)

	return e
}
