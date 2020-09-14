package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
