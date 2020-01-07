package main

import (
	"cryptocurrency_trading/app/controllers"
	"log"
)

func main() {
	// utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
