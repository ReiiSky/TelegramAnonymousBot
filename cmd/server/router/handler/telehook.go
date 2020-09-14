package handler

import (
	"github.com/labstack/echo/v4"
)

// Telehook handler for capturing data from telegram
func Telehook(c echo.Context) error {
	return c.NoContent(204)
}
