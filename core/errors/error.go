package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorFormat(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

func BadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}
