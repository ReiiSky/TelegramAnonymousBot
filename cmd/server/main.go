package main

import (
	"fmt"
	"os"

	"github.com/Satssuki/tele-anon-bot-queue/cmd/server/router"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("env file load error with message: %v", err.Error()))
	}
	pkg.SetDefaultTeleClient(os.Getenv("APIKEY"))
	router.SetupHandler()
	router.Router.Start(":6007")
}
