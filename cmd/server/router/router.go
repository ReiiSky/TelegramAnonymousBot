package router

import (
	"github.com/Satssuki/tele-anon-bot-queue/cmd/server/router/handler"
	"github.com/labstack/echo/v4"
)

// Router route register
var Router = echo.New()

// SetupHandler ..
func SetupHandler() {
	Router.POST("/telehook", handler.Telehook)
}
