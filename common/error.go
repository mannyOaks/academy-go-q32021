package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequestError(c echo.Context, message string) error {
	res := ErrorResponse{
		Message: message,
		Error:   nil,
	}
	return c.JSON(http.StatusBadRequest, res)
}

func NotFoundError(c echo.Context, id string) error {
	s := fmt.Sprintf("Movie %s not found", id)

	res := ErrorResponse{
		Message: s,
		Error:   nil,
	}
	return c.JSON(http.StatusNotFound, res)
}

func InternalServerError(c echo.Context, err error) error {
	res := ErrorResponse{
		Message: "Something wrong in server",
		Error:   err,
	}
	return c.JSON(http.StatusInternalServerError, res)
}
