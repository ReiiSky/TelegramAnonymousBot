package module

import (
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
)

// CommandList command contained list
var CommandList = []string{
	"/search",
	"/next",
	"/stop",
	"/stats",
}

// GetCommand ..
func GetCommand(message *pkg.Message) string {
	for _, y := range CommandList {
		if message.Content.Text == y {
			return y
		}
	}
	return ""
}
