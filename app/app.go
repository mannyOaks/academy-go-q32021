package app

import (
	"mrobles_app/controllers"
	"mrobles_app/infrastructure"
	"mrobles_app/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RunApp - Application startup and configuration
func RunApp() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] - ${method} ${status} ${uri} - [${remote_ip}]\n",
	}))

	handleMovies(e)
	e.Start(":5000")
}

func handleMovies(e *echo.Echo) {
	moviesService := services.NewMovieService(infrastructure.MovieRepo{})
	moviesController := controllers.NewMovieHandler(moviesService)

	e.GET("/movies/:id", moviesController.Controller)
}
