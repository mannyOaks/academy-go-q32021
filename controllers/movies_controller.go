package controllers

import (
	"net/http"
	"strconv"

	"mrobles_app/common"

	"github.com/labstack/echo/v4"
)

type filter interface {
	FindMovie(id string) (*common.Movie, error)
}

type MovieHandler struct {
	service filter
}

func NewMovieHandler(service filter) MovieHandler {
	return MovieHandler{service: service}
}

func (mv MovieHandler) Controller(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return common.BadRequestError(c, "Param {id} must be numeric")
	}

	movie, err := mv.service.FindMovie(id)
	if err != nil {
		return common.NotFoundError(c, id)
	}

	res := common.MovieResponse{
		Movie: *movie,
	}
	return c.JSON(http.StatusOK, res)
}
