package main

import (
	"github.com/joho/godotenv"
	"github.com/revogabe/go-jobsdev/config"
	"github.com/revogabe/go-jobsdev/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Initialize godotenv
	godotenv.Load()

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize the routes
	router.Initialize()
}
