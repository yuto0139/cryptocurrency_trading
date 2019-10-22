package main

import (
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
