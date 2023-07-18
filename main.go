package main

import (
	"github.com/RicardoArielSouza/gopportunies.git/config"
	"github.com/RicardoArielSouza/gopportunies.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
