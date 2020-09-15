package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Satssuki/tele-anon-bot-queue/cmd/server/router"
	"github.com/Satssuki/tele-anon-bot-queue/internal/lib"
	"github.com/Satssuki/tele-anon-bot-queue/internal/service"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
	"github.com/joho/godotenv"
)

func main() {
	// Set more worker to load more task
	runtime.GOMAXPROCS(8)

	// load environtment
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("env file load error with message: %v", err.Error()))
	}

	// Set Default package
	pkg.SetDefaultTeleClient(os.Getenv("APIKEY"))

	// Start worker to deliver message
	lib.StartWorker(3)

	// Start worker to scheduling the user matcher
	service.UserWorkerQueue()

	router.SetupHandler()
	router.Router.Start(":6007")
}
