package main

import (
	"cryptocurrency_trading/app/controllers"
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
