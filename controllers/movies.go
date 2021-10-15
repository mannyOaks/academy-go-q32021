package controllers

import (
	"net/http"
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/common"
	"github.com/mannyOaks/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type movieService interface {
	FindMovie(string) (entities.Movie, error)
	FindMovies(string, int, int) ([]entities.Movie, error)
}

// MovieController represents the controller the movie router uses
type MovieController struct {
	service movieService
}

// NewMovieController - receives a `movieService` and instantiates a Movie Controller
func NewMovieController(srv movieService) MovieController {
	return MovieController{srv}
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

// GetMovies - Returns all the movies saved in the csv file
func (mv MovieController) GetMovies(c echo.Context) error {
	typeParam := c.QueryParam("type")
	if typeParam == "" || (typeParam != "odd" && typeParam != "even") {
		return common.BadRequestError(c, "[type] param must be a value of \"odd\" or \"even\"")
	}

	items, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return common.BadRequestError(c, "[items] param must be an integer")
	}

	itemsPerWorker, err := strconv.Atoi(c.QueryParam("items_per_worker"))
	if err != nil {
		return common.BadRequestError(c, "[items_per_workers] param must be an integer")
	}

	if items < itemsPerWorker {
		return common.BadRequestError(c, "[items] param must be bigger than [items_per_worker] param")
	}

	movies, err := mv.service.FindMovies(typeParam, items, itemsPerWorker)
	if err != nil {
		return common.InternalServerError(c, err)
	}

	res := entities.GetMoviesResponse{Movies: movies}
	return c.JSON(200, res)
}
