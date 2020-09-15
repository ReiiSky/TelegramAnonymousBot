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
	userID := fmt.Sprint(msg.Content.From.ID)
	partnerID, isExist := service.ReadUserPartner(userID)
	if !isExist || len(partnerID) < 2 {
		if model.DefaultUserQueue().IsExist(userID) {
			return "You already call /search or /next command, and now you are in queue", nil
		}
		return "Currently you don't have any partner please call type /search command", nil
	}
	service.RemovePartner(userID, partnerID)
	return "You and your partner is unpaired", nil
}

// NextCommand ..
func NextCommand(msg *pkg.Message) (string, error) {
	userID := fmt.Sprint(msg.Content.From.ID)
	partnerID, isExist := service.ReadUserPartner(userID)

	if isExist {
		service.RemovePartner(userID, partnerID)
		model.DefaultUserQueue().Insert(userID)
		return "You skipped current partner", nil
	}
	return "You don't have a partner or still in queue", nil
}

// StatsCommand ..
func StatsCommand(msg *pkg.Message) (string, error) {
	return fmt.Sprintf("Paired count: %v", service.CountOfPartner()), nil
}
