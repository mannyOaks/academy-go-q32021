package controllers

import (
	"net/http"
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/common"
	"github.com/mannyOaks/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type movieService interface {
	FindMovie(id string) (entities.Movie, error)
}

// MovieController represents the controller the movie router uses
type MovieController struct {
	service movieService
}

// NewMovieController - receives a `movieService` and instantiates a Movie Controller
func NewMovieController(srv movieService) MovieController {
	return MovieController{service: srv}
}

// GetMovie - Returns a movie if found by the external API, also saves it in the csv file
func (mv MovieController) GetMovie(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return common.BadRequestError(c, "Param {id} must be numeric")
	}

	movie, err := mv.service.FindMovie(id)
	if err != nil {
		return common.InternalServerError(c, err)
	}

	if movie == (entities.Movie{}) {
		return common.NotFoundError(c, id)
	}

	res := entities.GetMovieResponse{
		Movie: movie,
	}
	return c.JSON(http.StatusOK, res)
}
