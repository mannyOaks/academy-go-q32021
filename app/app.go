package app

import (
	"mrobles_app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API configuration
func RunApp() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/movies", routes.GetMovies)
	e.GET("/movies/:id", routes.GetMovie)
	e.Start(":5000")
}
