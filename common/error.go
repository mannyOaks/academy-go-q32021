package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mannyOaks/academy-go-q32021/entities"
)

func BadRequestError(c echo.Context, message string) error {
	res := entities.ErrorResponse{
		Message: message,
	}
	return c.JSON(http.StatusBadRequest, res)
}

func NotFoundError(c echo.Context, id string) error {
	s := fmt.Sprintf("Movie %s not found", id)

	res := entities.ErrorResponse{
		Message: s,
	}
	return c.JSON(http.StatusNotFound, res)
}

func InternalServerError(c echo.Context, err error) error {
	res := entities.ErrorResponse{
		Message: "Something wrong in server",
	}
	return c.JSON(http.StatusInternalServerError, res)
}
