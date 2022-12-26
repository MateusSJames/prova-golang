package router

import (
	"Golang/src/app/api/endpoints/handlers"

	"github.com/labstack/echo/v4"
)

func loadAuthRoutes(api *echo.Group) {

	handler := handlers.NewAuthHandlers()
	api.POST("/verify", handler.SignUp)
}
