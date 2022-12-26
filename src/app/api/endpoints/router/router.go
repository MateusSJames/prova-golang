package router

import (
	"github.com/labstack/echo/v4"
)

func Start() *echo.Echo {
	router := echo.New()
	api := router.Group("/api")
	loadAuthRoutes(api)
	return router
}
