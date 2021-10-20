package controllers

import (
	"net/http"
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/common"
	"github.com/mannyOaks/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type MovieService interface {
	FindMovie(string) (*entities.Movie, error)
	FindMovies(string, int, int) ([]entities.Movie, error)
}

// MovieController represents the controller the movie router uses
type MovieController struct {
	service MovieService
}

// NewMovieController - receives a `MovieService` and instantiates a Movie Controller
func NewMovieController(srv MovieService) MovieController {
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

	if movie == nil {
		return common.NotFoundError(c, id)
	}

	res := entities.GetMovieResponse{
		Movie: *movie,
	}
	return c.JSON(http.StatusOK, res)
}

// GetMovies - Returns all the movies saved in the csv file
func (mv MovieController) GetMovies(c echo.Context) error {
	filter := c.QueryParam("type")
	if filter != "odd" && filter != "even" && filter != "" {
		return common.BadRequestError(c, "[type] param must be empty or one of ['odd', 'even']")
	}

	itemsParam := c.QueryParam("items")
	items, err := strconv.Atoi(itemsParam)
	if err != nil && itemsParam != "" {
		return common.BadRequestError(c, "[items] param must be an integer")
	}

	itemsPerWorker, err := strconv.Atoi(c.QueryParam("items_per_worker"))
	if err != nil {
		return common.BadRequestError(c, "[items_per_workers] param must be an integer")
	}

	if items > 0 && items < itemsPerWorker {
		return common.BadRequestError(c, "[items] param must be bigger than [items_per_worker] param")
	}

	movies, err := mv.service.FindMovies(filter, items, itemsPerWorker)
	if err != nil {
		return common.InternalServerError(c, err)
	}

	res := entities.GetMoviesResponse{Movies: movies}
	return c.JSON(200, res)
}
