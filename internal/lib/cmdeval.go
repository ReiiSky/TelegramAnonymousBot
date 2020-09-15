package lib

import (
	"fmt"

	"github.com/Satssuki/tele-anon-bot-queue/internal/domain/model"
	"github.com/Satssuki/tele-anon-bot-queue/internal/service"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
)

// CommandList i make this variable to reduce usage of if else
var CommandList = map[string]func(*pkg.Message) (string, error){
	"/search": SearchCommand,
	"/stop":   StopCommand,
	"/next":   NextCommand,
	"/stats":  StatsCommand,
}

// SearchCommand ..
func SearchCommand(msg *pkg.Message) (string, error) {
	userID := fmt.Sprint(msg.Content.From.ID)
	partnerID, isExist := service.ReadUserPartner(userID)
	if !isExist || len(partnerID) < 2 {
		if model.DefaultUserQueue().IsExist(userID) {
			return "You already added before to matching queue", nil
		}
		model.DefaultUserQueue().Insert(userID)
		return "You added to matching queue", nil
	}
	return "You already had a partner", nil
}

// StopCommand ..
func StopCommand(msg *pkg.Message) (string, error) {
	return "You stopped the chatting with current partner", nil
}

// NextCommand ..
func NextCommand(msg *pkg.Message) (string, error) {
	return "You skipped current partner", nil
}

// StatsCommand ..
func StatsCommand(msg *pkg.Message) (string, error) {
	return "heheheheh", nil
}
