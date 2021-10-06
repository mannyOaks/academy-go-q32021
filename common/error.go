package common

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func BadRequestError(c echo.Context, message string) error {
	log.Fatal(message)

	res := ErrorResponse{
		Message: message,
		Error:   nil,
	}
	data, _ := json.Marshal(res)
	return c.JSON(400, data)
}

func NotFoundError(c echo.Context, id int) error {
	s := fmt.Sprintf("Movie %d not found", id)
	log.Fatal(s)

	res := ErrorResponse{
		Message: s,
		Error:   nil,
	}
	data, _ := json.Marshal(res)
	return c.JSON(404, data)
}

func InternalServerError(c echo.Context, err error) error {
	log.Fatal(err)
	res := ErrorResponse{
		Message: "Something wrong in server",
		Error:   err,
	}
	data, _ := json.Marshal(res)
	return c.JSON(500, data)
}
