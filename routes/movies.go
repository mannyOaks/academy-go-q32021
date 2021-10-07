package routes

import (
	"net/http"
	"strconv"

	"mrobles_app/common"
	"mrobles_app/services"

	"github.com/labstack/echo/v4"
)

func GetMovies(c echo.Context) error {
	movies, err := services.FindMovies()
	if err != nil {
		return common.InternalServerError(c, err)
	}

	res := common.MoviesResponse{
		Movies: movies,
	}
	return c.JSON(http.StatusOK, res)
}

func GetMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return common.BadRequestError(c, "Param {id} must be numeric")
	}

	movie, err := services.FindMovie(id)
	if err != nil || movie == (common.Movie{}) {
		return common.NotFoundError(c, id)
	}

	res := common.MovieResponse{
		Movie: movie,
	}
	return c.JSON(http.StatusOK, res)
}
