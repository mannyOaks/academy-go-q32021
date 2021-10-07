package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequestError(c echo.Context, message string) error {

	res := ErrorResponse{
		Message: message,
		Error:   nil,
	}
	data, _ := json.Marshal(res)
	return c.JSON(http.StatusBadRequest, data)
}

func NotFoundError(c echo.Context, id int) error {
	s := fmt.Sprintf("Movie %d not found", id)

	res := ErrorResponse{
		Message: s,
		Error:   nil,
	}
	data, _ := json.Marshal(res)
	return c.JSON(http.StatusNotFound, data)
}

func InternalServerError(c echo.Context, err error) error {
	res := ErrorResponse{
		Message: "Something wrong in server",
		Error:   err,
	}
	data, _ := json.Marshal(res)
	return c.JSON(http.StatusInternalServerError, data)
}
