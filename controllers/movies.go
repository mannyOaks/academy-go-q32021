package controllers

import (
	"net/http"
	"strconv"

	"mrobles_app/common"

	"github.com/labstack/echo/v4"
)

type service interface {
	FindMovie(id string) (common.Movie, error)
}

type MovieHandler struct {
	srv service
}

func NewMovieHandler(srv service) MovieHandler {
	return MovieHandler{srv: srv}
}

func (mv MovieHandler) Controller(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return common.BadRequestError(c, "Param {id} must be numeric")
	}

	movie, err := mv.srv.FindMovie(id)
	if err != nil {
		return common.InternalServerError(c, err)
	}

	if movie == (common.Movie{}) {
		return common.NotFoundError(c, id)
	}

	res := common.MovieResponse{
		Movie: movie,
	}
	return c.JSON(http.StatusOK, res)
}
