package handler

import (
	"fmt"

	"github.com/Satssuki/tele-anon-bot-queue/internal/lib"
	"github.com/Satssuki/tele-anon-bot-queue/internal/module"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
	"github.com/labstack/echo/v4"
)

// Telehook handler for capturing data from telegram
func Telehook(c echo.Context) error {
	capturedMessage := &pkg.Message{}
	err := pkg.UnmarshalByIOReader(c.Request().Body, &capturedMessage)
	if err != nil {
		return c.NoContent(400)
	}
	command := module.GetCommand(capturedMessage)
	if len(command) > 1 {
		action := lib.CommandList[command]
		message, err := action(capturedMessage)
		if err != nil {
			return c.NoContent(501)
		}
		compiledMessage := pkg.
			GetDefaultTeleClient().
			TextMessageBuilder().
			ChatID(fmt.Sprint(capturedMessage.Content.From.ID)).
			Content(message).
			ReplyTo(fmt.Sprint(capturedMessage.Content.MessageID))

		pkg.GetDefaultTeleClient().Push(compiledMessage)
		return c.NoContent(204)
	}
	lib.MessageDeliverQueue <- capturedMessage
	return c.NoContent(201)
}
