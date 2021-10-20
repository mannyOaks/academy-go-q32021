package app

import (
	"os"

	"github.com/mannyOaks/academy-go-q32021/controllers"
	"github.com/mannyOaks/academy-go-q32021/infrastructure"
	"github.com/mannyOaks/academy-go-q32021/services"

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
	e.Start(":" + os.Getenv("PORT"))
}

func handleMovies(e *echo.Echo) {
	movieRepo := infrastructure.NewMovieRepo()
	workerPool := infrastructure.NewMovieWorkerPool()
	moviesService := services.NewMovieService(movieRepo, workerPool)
	moviesController := controllers.NewMovieController(moviesService)

	e.GET("/movies", moviesController.GetMovies)
	e.GET("/movies/:id", moviesController.GetMovie)
}
