package main

import (
	"cryptocurrency_trading/app/controllers"
	"cryptocurrency_trading/app/models"
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
}
