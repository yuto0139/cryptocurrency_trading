package main

import (
	"cryptocurrency_trading/app/controllers"
	"cryptocurrency_trading/app/models"
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	df, _ := models.GetAllCandle(config.Config.ProductCode, time.Minute, 365)
	fmt.Printf("%+v\n", df.OptimizeParams())
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
