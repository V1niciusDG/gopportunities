package main

import (
	"github.com/V1niciusDG/gopportunities/config"
	"github.com/V1niciusDG/gopportunities/router"
)

var (
	logger *config.Logger
)


func main() {
logger = config.GetLogger("main")
config.GetExampleEnv()
	// initialize configs
	err := config.Init()
	if err != nil{
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//chama router
	router.Initialize()
}