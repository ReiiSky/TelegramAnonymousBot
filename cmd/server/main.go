package main

import (
	"github.com/Satssuki/tele-anon-bot-queue/cmd/server/router"
)

func main() {
	router.SetupHandler()
	router.Router.Start(":6007")
}
