package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequestError(c echo.Context, message string) error {
	res := errorResponse{
		Message: message,
	}
	return c.JSON(http.StatusBadRequest, res)
}

func NotFoundError(c echo.Context, id string) error {
	s := fmt.Sprintf("Movie %s not found", id)

	res := errorResponse{
		Message: s,
	}
	return c.JSON(http.StatusNotFound, res)
}

func InternalServerError(c echo.Context, err error) error {
	res := errorResponse{
		Message: "Something wrong in server",
	}
	return c.JSON(http.StatusInternalServerError, res)
}
