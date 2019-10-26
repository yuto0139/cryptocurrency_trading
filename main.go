package main

import (
	"cryptocurrency_trading/bitflyer"
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
