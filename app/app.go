package app

import (
	"mrobles_app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API configuration
func RunApp() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] - ${method} ${status} ${uri} - [${remote_ip}]\n",
	}))

	e.GET("/movies", routes.GetMovies)
	e.GET("/movies/:id", routes.GetMovie)
	e.Start(":5000")
}

// `${method} ${originalUrl} ${statusCode} ${contentLength} - ${userAgent} ${ip}`,
